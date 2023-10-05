package mq

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"time"
)

type Kafka struct {
	brokers []string
	group   string
	tls     *tls.Config

	stopCtx context.Context
	l       *logger.Logger
}

// New - создаем клиента для брокера сообщений
func New(ctx context.Context, brokers []string, t *tls.Config, timeout time.Duration, l *logger.Logger) *Kafka {
	k := new(Kafka)
	k.l = l
	k.brokers = brokers
	k.stopCtx = ctx
	k.tls = t
	timeoutCtx, cancel := context.WithTimeout(k.stopCtx, timeout)
	defer cancel()

	// проверяем соединения с брокерами
	if err := connectionCheck(timeoutCtx, brokers); err != nil {
		k.l.Debugw(err.Error(), "brokers", brokers, "timeout", timeout)
		k.l.Panic(err)
	}

	//go func() {
	//	for {
	//		k.l.Info("check conn")
	//		_, err := kafka.Dial("tcp", k.brokers[0])
	//		if err != nil {
	//			k.l.Error(err)
	//		}
	//		time.Sleep(1 * time.Second)
	//	}
	//}()

	return k
}

// Subscribe - подписываемся на топик
// стартует воркер, который слушает сообщения из кафки
func (k *Kafka) subscribe(ctx context.Context, topic string, handler func(m []byte) error) {
	di := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  k.brokers,
		GroupID:  uuid.New().String(),
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
		Dialer:   di,
	})

	for {
		select {
		case <-ctx.Done():
			return
		default:
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				k.l.Error(err)
			}
			k.l.Debug(m.Topic, string(m.Value))

			if len(m.Value) > 0 {
				if err := handler(m.Value); err != nil {
					k.l.Error(err)
				}
			}
		}
	}
}

func (k *Kafka) publish(topic string, v interface{}) (err error) {
	value, err := json.Marshal(v)
	if err != nil {
		return
	}

	k.l.Debug(string(value))

	w := &kafka.Writer{
		Addr:                   kafka.TCP(k.brokers...),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	if k.tls != nil {
		w.Transport = &kafka.Transport{TLS: k.tls}
	}

	msg := kafka.Message{
		Value: value,
		Time:  time.Now(),
	}

	err = w.WriteMessages(context.Background(), msg)
	if err != nil {
		err = fmt.Errorf("write message err: %w", err)
		return
	}

	err = w.Close()
	if err != nil {
		err = fmt.Errorf("close err: %w", err)
	}
	return
}

// connectionCheck - запускаем воркер проверять доступность брокеров
// завершиться, если все переданные broker'ы доступны, либо по тайауту
func connectionCheck(ctx context.Context, brokers []string) error {
	m := make(map[string]*kafka.Conn)
	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout")
		default:
			// проверяем подключение ко всем брокерам
			// будем пытаться подключиться, пока не подключимся
			// todo подключаться повторно только к тем брокерам, которые еще не подключены
			for _, broker := range brokers {
				// если брокер уже подключен, то пропускаем его
				_, ok := m[broker]
				if ok {
					continue
				}
				conn, err := kafka.Dial("tcp", broker)
				if err == nil {
					m[broker] = conn
				}
				bl, err := conn.Brokers()
				fmt.Println(bl)
			}
			// если ко всем брокерам подключились, выходим из цикла
			if len(m) == len(brokers) {
				return nil
			}
			// задержка между попытками
			time.Sleep(time.Second * 10)
		}
	}
}

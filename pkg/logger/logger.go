package logger

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Encoding string

const (
	Json    Encoding = "json"
	Console Encoding = "console"
)

type Logger struct {
	level    Level
	zap      *zap.SugaredLogger
	config   *zap.Config
	baseOpts []zap.Option
}
type Level zapcore.Level

type Zap struct {
	sugarClient *zap.SugaredLogger
	client      *zap.Logger
}

var l = &Logger{}

// New инициализация логгера, передаем в режиме дебага или нет
// encoding - json | console
func New(debug, showTime bool, encoding Encoding) *Logger {
	l := new(Logger)
	var err error

	// setup logs
	lvl := "info"
	isDev := false
	disableStack := true

	if debug {
		lvl = "debug"
		isDev = true
		disableStack = false
	}

	// setup encoding
	var encodingStr string
	encConf := zapcore.EncoderConfig{
		FunctionKey:    zapcore.OmitKey,
		EncodeTime:     stampTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   callerEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
	}

	if showTime {
		encConf.TimeKey = "ts"
	}

	switch encoding {
	case Json:
		encodingStr = "json"
	case Console:
		encodingStr = "console"
		encConf.ConsoleSeparator = " "
	default:
		fmt.Println("Logger init error: invalid encoding - 'console' or 'json'")
		return nil
	}

	l.config = &zap.Config{
		Level:             levelToAtomic(parseLevel(lvl)),
		Development:       isDev,
		DisableCaller:     false,
		DisableStacktrace: disableStack,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:      encodingStr,
		EncoderConfig: encConf,
		//OutputPaths:      []string{"/var/log/syslog"},
		//ErrorOutputPaths: []string{"/var/log/syslog"},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	l.baseOpts = append(l.baseOpts, zap.AddCallerSkip(1))
	l.level = parseLevel(lvl)
	err = l.build()
	if err != nil {
		fmt.Println("Logger init error: ", err)
		return nil
	}
	return l
}

func (l *Logger) WithField(field string) {
	l.build(zap.Fields(zap.Field{
		Key:  field,
		Type: zapcore.StringType,
	}))
}

func (l *Logger) build(opts ...zap.Option) (err error) {
	opts = append(l.baseOpts, opts...)

	lg, err := l.config.Build(opts...)
	if err != nil {
		l.Error(errors.Wrap(err, "Logger init error"))
		return
	}
	l.zap = lg.Sugar()
	return
}

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	//arr := strings.Split(caller.Function, ".")
	//funName := arr[len(arr)-1]
	enc.AppendString(caller.TrimmedPath())
}

func stampTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	format := "Jan 02 15:04:05"
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, format)
		return
	}

	enc.AppendString(t.Format(format))
}

func levelToAtomic(lvl Level) zap.AtomicLevel {
	return zap.NewAtomicLevelAt(zapcore.Level(lvl))
}

func parseLevel(lvl string) (level Level) {
	switch lvl {
	case "debug":
		level = Level(zap.DebugLevel)
	case "info":
		level = Level(zap.InfoLevel)
	case "warning":
		level = Level(zap.WarnLevel)
	case "error":
		level = Level(zap.ErrorLevel)
	case "panic":
		level = Level(zap.PanicLevel)
	case "fatal":
		level = Level(zap.FatalLevel)
	}
	return
}

// Fatal followed by a call to os.Exit(1).
func (l *Logger) Fatal(msg ...interface{}) {
	l.zap.Fatal(msg...)
}

// Panic followed by a call to panic().
func (l *Logger) Panic(msg ...interface{}) {
	l.zap.Panic(msg...)
}

// Error logs a message using ERROR as log level.
func (l *Logger) Error(msg ...interface{}) {
	l.zap.Error(msg...)
}

// Warning logs a message using WARNING as log level.
func (l *Logger) Warning(msg ...interface{}) {
	l.zap.Warn(msg...)
}

// Info logs a message using INFO as log level.
func (l *Logger) Info(msg ...interface{}) {
	l.zap.Info(msg...)
}

// Debug logs a message using DEBUG as log level.
func (l *Logger) Debug(msg ...interface{}) {
	l.zap.Debug(msg...)
}

// Fatalf followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.zap.Fatalf(format, args...)
}

// Panicf followed by a call to panic().
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.zap.Panicf(format, args...)
}

// Errorf logs a message using ERROR as log level.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.zap.Errorf(format, args...)
}

// Warningf logs a message using WARNING as log level.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.zap.Warnf(format, args...)
}

// Infof logs a message using INFO as log level.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.zap.Infof(format, args...)
}

// Debugf logs a message using DEBUG as log level.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.zap.Debugf(format, args...)
}

// Fatalw followed by a call to os.Exit(1).
func (l *Logger) Fatalw(msg string, keyAndValues ...interface{}) {
	l.zap.Fatalw(msg, keyAndValues...)
}

// Errorw logs a message using ERROR as log level.
func (l *Logger) Errorw(msg string, keyAndValues ...interface{}) {
	l.zap.Errorw(msg, keyAndValues...)
}

// Warningw logs a message using WARNING as log level.
func (l *Logger) Warningw(msg string, keyAndValues ...interface{}) {
	l.zap.Warnw(msg, keyAndValues...)
}

// Infow logs a message using INFO as log level.
func (l *Logger) Infow(msg string, keyAndValues ...interface{}) {
	l.zap.Infow(msg, keyAndValues...)
}

// Debugw logs a message using DEBUG as log level.
func (l *Logger) Debugw(msg string, keyAndValues ...interface{}) {
	l.zap.Debugw(msg, keyAndValues...)
}

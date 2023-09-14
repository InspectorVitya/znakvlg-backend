package main

import (
	"context"
	"errors"
	"github.com/InspectorVitya/znakvlg-backend/internal/app"
	"github.com/InspectorVitya/znakvlg-backend/internal/config"
	"github.com/InspectorVitya/znakvlg-backend/internal/database"
	"github.com/InspectorVitya/znakvlg-backend/internal/transport/rest"
	"github.com/InspectorVitya/znakvlg-backend/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	log := logger.New(true, true, logger.Console)
	cfg := config.LoadConfig()

	db, err := database.New(context.TODO(), cfg.DataBase.DBURL)
	if err != nil {
		log.Fatalf("cannot init DB: %w", err)
	}
	service, err := app.New(log, db)
	if err != nil {
		log.Fatalf("cannot init app: %w", err)
	}

	httpServer := rest.New(service, cfg.HTTP.Port, log)

	go func() {
		if err := httpServer.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server")
		}
	}()

	log.Info("Server started")
	log.Info("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit

	const timeout = 10 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := httpServer.Stop(ctx); err != nil {
		log.Error(err)
	}
	if err := db.Close(ctx); err != nil {
		log.Error(err)
	}

}

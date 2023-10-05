package main

import (
	"context"
	"errors"
	"github.com/InspectorVitya/znakvlg-backend/internal/app"
	"github.com/InspectorVitya/znakvlg-backend/internal/config"
	"github.com/InspectorVitya/znakvlg-backend/internal/repository"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	"github.com/InspectorVitya/znakvlg-backend/internal/transport/rest"
	manager "github.com/InspectorVitya/znakvlg-backend/pkg/jwt"
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

	repo := repository.New()

	store, err := storage.New(context.TODO(), cfg.DBURL)
	if err != nil {
		log.Fatalf("cannot connect db: %w", err)
	}

	companyApp, err := app.NewCompany(log, repo, store)
	if err != nil {
		log.Fatalf("cannot init company app: %w", err)
	}

	userApp, err := app.NewUser(log, repo, store)
	if err != nil {
		log.Fatalf("cannot init company app: %w", err)
	}

	authManager, err := manager.JWTAuthService(cfg.JwtSecret)
	if err != nil {
		log.Fatalf("cannot init authManager: %w", err)
	}

	authApp, err := app.NewAuth(log, repo, authManager, store)
	if err != nil {
		log.Fatalf("cannot init company app: %w", err)
	}

	httpServer := rest.New(companyApp, userApp, authApp, cfg.HTTP.Port, log)

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
	if err := store.Close(); err != nil {
		log.Error(err)
	}

}

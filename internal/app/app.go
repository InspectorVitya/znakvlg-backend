package app

import (
	"github.com/InspectorVitya/znakvlg-backend/pkg/logger"
)

type App struct {
	db DB
	l  *logger.Logger
}

func New(logger *logger.Logger, db DB) (*App, error) {
	app := &App{
		db: db,
		l:  logger,
	}
	return app, nil
}

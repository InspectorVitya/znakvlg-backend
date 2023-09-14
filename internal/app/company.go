package app

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
)

func (app *App) CreateCompany(ctx context.Context, company model.Company, platesUse model.PlatesUse) error {
	store, err := app.createStore(ctx, platesUse)
	if err != nil {
		return err
	}
	err = app.db.InitCompany(ctx, company, store)
	return err
}

func (app *App) GetCompany(ctx context.Context, companyID uint32) (model.Company, error) {
	return app.db.GetCompanyByID(ctx, companyID)
}

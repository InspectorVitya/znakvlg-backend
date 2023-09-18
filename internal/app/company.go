package app

import (
	"context"
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	"github.com/InspectorVitya/znakvlg-backend/pkg/logger"
)

type Company struct {
	db      CompanyDB
	storage storage.Storage
	l       *logger.Logger
}

func NewCompany(logger *logger.Logger, db CompanyDB, storage storage.Storage) (*Company, error) {
	app := &Company{
		db:      db,
		l:       logger,
		storage: storage,
	}
	return app, nil
}

func (c *Company) CreateCompany(ctx context.Context, company model.Company, platesUse model.PlatesUse) error {
	tx, err := c.storage.BeginTx(ctx, "create company")
	if err != nil {
		return err
	}
	defer func() {
		err = tx.CommitRollback(ctx, err)
	}()

	companyID, err := c.db.InsertCompany(ctx, tx, company)
	if err != nil {
		err = fmt.Errorf("create company err: %w", err)
		return err
	}

	err = c.createStoreForCompany(ctx, tx, companyID, platesUse)
	if err != nil {
		err = fmt.Errorf("create company store err: %w", err)
		return err
	}
	return err
}

// todo
func (c *Company) createStoreForCompany(ctx context.Context, tx storage.StorageTx, companyID uint32, platesUse model.PlatesUse) error {
	store := make([]model.Store, 0, len(platesUse))
	for i := range platesUse {
		store = append(store, model.Store{
			CompanyID:   companyID,
			PlateTypeID: platesUse[i],
			Used:        true,
		})
	}

	err := c.db.InsertStore(ctx, tx, store)
	return err
}

func (c *Company) GetCompanyByID(ctx context.Context, companyID uint32) (model.Company, error) {
	return c.db.SelectCompanyByID(ctx, c.storage, companyID)
}

package database

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
)

func (db *DB) SelectPlateTypes(ctx context.Context) (model.PlateTypes, error) {
	return model.PlateTypes{}, nil
}

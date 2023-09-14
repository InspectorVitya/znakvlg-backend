package database

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (db *DB) insertStore(ctx context.Context, tx pgx.Tx, companyID uint32, store []model.Store) error {
	queryStore := sq.Insert("store").Columns("company_id", "plate_type_id", "used").PlaceholderFormat(sq.Dollar)

	for i := range store {
		queryStore = queryStore.Values(companyID, store[i].PlateTypeID, store[i].Used)
	}

	sqlInsert, args, err := queryStore.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sqlInsert, args...)
	return err
}

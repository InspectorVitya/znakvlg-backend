package repository

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	sq "github.com/Masterminds/squirrel"
)

func (r Repository) InsertStore(ctx context.Context, q storage.Query, store []model.Store) error {
	queryStore := sq.Insert("store").Columns("company_id", "plate_type_id", "used").PlaceholderFormat(sq.Dollar)

	for i := range store {
		queryStore = queryStore.Values(store[i].CompanyID, store[i].PlateTypeID, store[i].Used)
	}

	sqlInsert, args, err := queryStore.ToSql()
	if err != nil {
		return err
	}
	err = q.Exec(ctx, "insert store", sqlInsert, args...)
	return err
}

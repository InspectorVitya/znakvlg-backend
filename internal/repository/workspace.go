package repository

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	sq "github.com/Masterminds/squirrel"
)

func (r Repository) InsertUserWorkspace(ctx context.Context, q storage.Query, userID string, WorkPlace model.WorkPlace) error {

	queryCompanyId := sq.Insert("work_place").Columns("company_id", "user_id").PlaceholderFormat(sq.Dollar)

	for i := range WorkPlace {
		queryCompanyId = queryCompanyId.Values(WorkPlace[i], userID)
	}
	sqlInsert, args, err := queryCompanyId.ToSql()
	if err != nil {
		return err
	}

	err = q.Exec(ctx, "insert user workspace", sqlInsert, args...)
	if err != nil {
		return err
	}
	return nil
}

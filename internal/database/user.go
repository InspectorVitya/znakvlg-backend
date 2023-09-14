package database

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *DB) InsertUser(ctx context.Context, user *model.Users) error {
	tx, err := db.conn.Begin(ctx)
	if err != nil {
		return err
	}

	user.ID = uuid.New().String()

	query, args, err := sq.Insert("users").
		Columns("id", "login", "email", "password", "blocked", "role_id", "comment", "phone_number", "name", "patronymic", "surname", "address").
		Values(user.ID, user.Login, user.Email, user.Password, user.Blocked, user.RoleID, user.Comment, user.PhoneNumber, user.Name, user.Patronymic, user.SurName, user.Address).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	if len(user.WorkPlace) > 0 {
		err := db.insertUserWorkspace(ctx, tx, user.ID, user.WorkPlace)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) insertUserWorkspace(ctx context.Context, tx pgx.Tx, userID string, WorkPlace model.WorkPlace) error {

	queryCompanyId := sq.Insert("work_place").Columns("company_id", "user_id").PlaceholderFormat(sq.Dollar)

	for i := range WorkPlace {
		queryCompanyId = queryCompanyId.Values(WorkPlace[i], userID)
	}
	sqlInsert, args, err := queryCompanyId.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, sqlInsert, args...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CheckLogin(ctx context.Context, login string) (bool, error) {

	query, args, err := sq.Select("login").
		Prefix("SELECT EXISTS (").
		From("users").
		Where("login=$1", login).
		Suffix(")").
		ToSql()
	if err != nil {
		return false, err
	}
	var exists bool
	err = db.conn.QueryRow(ctx, query, args...).Scan(&exists)

	return exists, err
}

func (db *DB) GetUserByID(ctx context.Context, id string) (model.Users, error) {
	sql, args, err := sq.Select("id", "login", "email", "password", "blocked", "comment", "phone_number", "surname", "name", "patronymic", "address", "last_login", "role_id").
		From("users").Where("id=$1", id).ToSql()
	if err != nil {
		return model.Users{}, err
	}

	user := model.Users{}
	err = db.conn.QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.Blocked, &user.Comment, &user.PhoneNumber, &user.SurName, &user.Name, &user.Patronymic, &user.Address, &user.LastLogin, &user.RoleID)
	if err != nil {
		return model.Users{}, err
	}

	if user.RoleID != 1 {
		ids, err := db.getCompanyIDs(ctx, user.ID)
		if err != nil {
			return model.Users{}, err
		}

		if len(ids) > 0 {
			user.WorkPlace = ids
		}
	}

	return user, nil
}

func (db *DB) getCompanyIDs(ctx context.Context, id string) (model.WorkPlace, error) {
	sql, args, err := sq.Select("company_id").From("work_place").Where("user_id=$1", id).ToSql()
	var ids model.WorkPlace

	rows, err := db.conn.Query(ctx, sql, args...)
	if err != nil {
		return model.WorkPlace{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint32
		err = rows.Scan(&id)
		if err != nil {
			return model.WorkPlace{}, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

package repository

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r Repository) InsertUser(ctx context.Context, q storage.Query, user *model.Users) (string, error) {

	user.ID = uuid.New().String()

	query, args, err := sq.Insert("users").
		Columns("id", "login", "email", "password", "blocked", "role_id", "comment", "phone_number", "name", "patronymic", "surname", "address").
		Values(user.ID, user.Login, user.Email, user.Password, user.Blocked, user.RoleID, user.Comment, user.PhoneNumber, user.Name, user.Patronymic, user.SurName, user.Address).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return user.ID, err
	}

	err = q.Exec(ctx, "insert user", query, args...)
	if err != nil {
		return user.ID, err
	}

	return user.ID, nil
}

func (r Repository) CheckLogin(ctx context.Context, q storage.Query, login string) (bool, error) {

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
	err = q.QueryOne(ctx, &exists, "check login", query, args...)

	return exists, err
}

func (r Repository) SelectUserByID(ctx context.Context, q storage.Query, id string) (model.Users, error) {
	sql, args, err := sq.Select("id", "login", "email", "password", "blocked", "comment", "phone_number", "surname", "name", "patronymic", "address", "last_login", "role_id").
		From("users").Where("id=$1", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return model.Users{}, err
	}

	user := model.Users{}
	err = q.QueryOne(ctx, &user, "select user by id", sql, args...)
	if err != nil {
		return model.Users{}, err
	}

	//todo изменить на один запрос и смапить
	if user.RoleID != 1 {
		ids, err := r.getUserWorkSpace(ctx, q, user.ID)
		if err != nil {
			return model.Users{}, err
		}

		if len(ids) > 0 {
			user.WorkPlace = ids
		}
	}

	return user, nil
}

func (r Repository) getUserWorkSpace(ctx context.Context, q storage.Query, id string) (model.WorkPlace, error) {
	sql, args, err := sq.Select("company_id").From("work_place").Where("user_id=$1", id).ToSql()
	if err != nil {
		return nil, err
	}
	var ids model.WorkPlace

	rows, err := q.QueryAll(ctx, "select work space", sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint32
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (r Repository) SelectUserByLogin(ctx context.Context, q storage.Query, login string) (model.Users, error) {
	sql, args, err := sq.Select("id", "login", "email", "password", "blocked", "comment", "phone_number", "surname", "name", "patronymic", "address", "last_login", "role_id").
		From("users").Where("login=$1", login).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return model.Users{}, err
	}

	user := model.Users{}
	err = q.QueryOne(ctx, &user, "select user by id", sql, args...)
	if err != nil {
		return model.Users{}, err
	}

	//todo изменить на один запрос и смапить
	if user.RoleID != 1 {
		ids, err := r.getUserWorkSpace(ctx, q, user.ID)
		if err != nil {
			return model.Users{}, err
		}

		if len(ids) > 0 {
			user.WorkPlace = ids
		}
	}

	return user, nil
}

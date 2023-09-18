package app

import (
	"context"
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	hasher "github.com/InspectorVitya/znakvlg-backend/pkg/bcrypt"
	"github.com/InspectorVitya/znakvlg-backend/pkg/logger"
)

type User struct {
	db      UserDB
	storage storage.Storage
	l       *logger.Logger
}

func NewUser(logger *logger.Logger, db UserDB, storage storage.Storage) (*User, error) {
	app := &User{
		db:      db,
		l:       logger,
		storage: storage,
	}
	return app, nil
}

func (u *User) CreateUser(ctx context.Context, user *model.Users) error {
	hashedPwd, err := hasher.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("hash password: %v", err)
	}
	user.Password = hashedPwd

	tx, err := u.storage.BeginTx(ctx, "create user")
	if err != nil {
		return err
	}
	defer func() {
		err = tx.CommitRollback(ctx, err)
	}()

	userID, err := u.db.InsertUser(ctx, tx, user)

	if len(user.WorkPlace) > 0 {
		err := u.db.InsertUserWorkspace(ctx, tx, userID, user.WorkPlace)
		if err != nil {
			return err
		}
	}

	return err

}

func (u *User) ValidateUser(ctx context.Context, user *model.Users) map[string]string {
	invalid := user.Validate()
	if len(invalid) > 0 {
		return invalid
	}

	exist, err := u.db.CheckLogin(ctx, u.storage, user.Login)
	if err != nil {
		invalid["error"] = err.Error()
		return invalid
	}

	if exist {
		invalid["login"] = "Такой логин существует"
		return invalid
	}

	return invalid
}

func (u *User) GetUserByID(ctx context.Context, userID string) (model.Users, error) {
	return u.db.SelectUserByID(ctx, u.storage, userID)
}

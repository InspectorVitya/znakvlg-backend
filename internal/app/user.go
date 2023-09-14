package app

import (
	"context"
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	hasher "github.com/InspectorVitya/znakvlg-backend/pkg/bcrypt"
)

func (app *App) CreateUser(ctx context.Context, user *model.Users) error {
	hashedPwd, err := hasher.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("hash password: %v", err)
	}
	user.Password = hashedPwd
	err = app.db.InsertUser(ctx, user)
	return err
}

func (app *App) ValidateUser(ctx context.Context, user *model.Users) map[string]string {
	invalid := user.Validate()
	if len(invalid) > 0 {
		return invalid
	}
	exist, err := app.db.CheckLogin(ctx, user.Login)
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

func (app *App) GetUserByID(ctx context.Context, userID string) (model.Users, error) {
	return app.db.GetUserByID(ctx, userID)
}

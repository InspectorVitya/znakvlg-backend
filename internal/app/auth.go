package app

import (
	"context"
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/model/consts"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	hasher "github.com/InspectorVitya/znakvlg-backend/pkg/bcrypt"
	"github.com/InspectorVitya/znakvlg-backend/pkg/logger"
	"github.com/google/uuid"
	"strings"
)

type Auth struct {
	db          UserAuthDB
	storage     storage.Storage
	authManager AuthManager
	l           *logger.Logger
}

func NewAuth(logger *logger.Logger, db UserAuthDB, authManager AuthManager, storage storage.Storage) (*Auth, error) {
	app := &Auth{
		db:          db,
		l:           logger,
		authManager: authManager,
		storage:     storage,
	}
	return app, nil
}

func (a *Auth) SignIn(ctx context.Context, userAuth model.RequestAuth) (token, refresh string, err error) {
	defer func() {
		if r := recover(); r != nil {
			a.l.Errorf("recovered in refresh token: %w ", r)
		}
	}()

	user, err := a.db.SelectUserByLogin(ctx, a.storage, strings.ToLower(userAuth.Login))
	if err != nil {
		return "", "", err
	}
	if !hasher.CheckPasswordHash(userAuth.Password, user.Password) {
		return "", "", model.ErrInvalidPassword
	}
	if user.Blocked {
		return "", "", model.ErrUserBlocked
	}
	token, err = a.authManager.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		a.l.Errorf("generate token: %w", err)
		return
	}
	//todo db insert refresh metainfo
	refresh = uuid.New().String()

	return token, refresh, err
}

func (a *Auth) AuthAdmin(_ context.Context, tokenString string) (model.JWTPayload, error) {
	claim, err := a.authManager.ValidateToken(tokenString)
	if err != nil {
		a.l.Errorf("auth admin: %v", err)
		return model.JWTPayload{}, err
	}
	if claim.UserRole != consts.AdminRole {
		err = model.ErrInvalidRole
		return model.JWTPayload{}, fmt.Errorf("user %s err: %w", claim.UserId, err)
	}
	// todo
	return model.JWTPayload{
		UserID:   claim.UserId,
		UserRole: claim.UserRole,
	}, nil
}

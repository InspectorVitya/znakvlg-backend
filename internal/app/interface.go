package app

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
)

//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g
type CompanyDB interface {
	SelectCompanyByID(ctx context.Context, q storage.Query, id uint32) (model.Company, error)
	InsertCompany(ctx context.Context, q storage.Query, company model.Company) (uint32, error)
	SelectCompanies(ctx context.Context, q storage.Query, ids []uint32) ([]*model.CompanyInfo, error)
	InsertStore(ctx context.Context, q storage.Query, store []model.Store) error
}

type UserDB interface {
	InsertUser(ctx context.Context, q storage.Query, user *model.Users) (string, error)
	InsertUserWorkspace(ctx context.Context, q storage.Query, userID string, WorkPlace model.WorkPlace) error
	CheckLogin(ctx context.Context, q storage.Query, login string) (bool, error)
	SelectUserByID(ctx context.Context, q storage.Query, id string) (model.Users, error)
}

type UserAuthDB interface {
	SelectUserByLogin(ctx context.Context, q storage.Query, login string) (model.Users, error)
	//todo refresh token
}

type AuthManager interface {
	GenerateToken(userID string, userRole uint8) (string, error)
	ValidateToken(encodedToken string) (*model.AuthClaims, error)
}

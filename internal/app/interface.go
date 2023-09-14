package app

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
)

type DB interface {
	UserDB
	CompanyDB
}

type CompanyDB interface {
	GetCompanyByID(ctx context.Context, id uint32) (model.Company, error)
	InitCompany(ctx context.Context, company model.Company, store []model.Store) error
	SelectCompanies(ctx context.Context, ids []uint32) ([]*model.CompanyInfo, error)
}

type UserDB interface {
	InsertUser(ctx context.Context, user *model.Users) error
	CheckLogin(ctx context.Context, login string) (bool, error)
	GetUserByID(ctx context.Context, id string) (model.Users, error)
}

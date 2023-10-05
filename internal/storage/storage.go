package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Query interface {
	Exec(ctx context.Context, queryName, query string, args ...any) error
	ExecWithResult(ctx context.Context, queryName, query string, args ...any) (any, error)
	QueryOne(ctx context.Context, dest any, queryName, query string, args ...any) error
	QueryAll(ctx context.Context, queryName, query string, args ...any) (*sqlx.Rows, error)
}

type Storage interface {
	Query
	Close() error
	BeginTx(ctx context.Context, txName string) (StorageTx, error)
}

type StorageTx interface {
	Query
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CommitRollback(ctx context.Context, err error) error
}

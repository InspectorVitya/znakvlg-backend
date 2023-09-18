package storage

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PgStorage struct {
	conn *sqlx.DB
}

func New(ctx context.Context, connURL string) (*PgStorage, error) {
	conn, err := sqlx.ConnectContext(ctx, "pgx", connURL)
	if err != nil {
		return nil, err
	}

	return &PgStorage{conn: conn}, nil
}

func (s *PgStorage) Close() error {
	return s.conn.Close()
}

func (s *PgStorage) Exec(ctx context.Context, queryName, query string, args ...any) error {
	_, err := s.conn.ExecContext(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("storage failed query %s: exec err: %w", queryName, err)
		return err
	}
	return nil
}

func (s *PgStorage) ExecWithResult(ctx context.Context, queryName, query string, args ...any) (any, error) {
	result, err := s.conn.ExecContext(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("storage failed query %s: exec with result err: %w", queryName, err)
		return nil, err
	}
	return result, nil
}

func (s *PgStorage) QueryOne(ctx context.Context, dest any, queryName, query string, args ...any) error {
	err := s.conn.GetContext(ctx, dest, query, args...)
	if err != nil {
		err = fmt.Errorf("storage failed query %s: query one err: %w", queryName, err)
		return err
	}
	return nil
}

func (s *PgStorage) QueryAll(ctx context.Context, dest any, queryName, query string, args ...any) error {
	err := s.conn.SelectContext(ctx, dest, query, args...)
	if err != nil {
		err = fmt.Errorf("storage failed query %s: query all err: %w", queryName, err)
		return err
	}

	return nil
}

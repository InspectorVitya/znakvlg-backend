package storage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type pgTx struct {
	tx     *sqlx.Tx
	txName string
}

func (s *PgStorage) BeginTx(ctx context.Context, txName string) (StorageTx, error) {
	tx, err := s.conn.BeginTxx(ctx, nil)
	if err != nil {
		err = fmt.Errorf("cannot begin tx %s err: %w", txName, err)
		return nil, err
	}
	return &pgTx{tx: tx, txName: txName}, nil
}

func (s *pgTx) Exec(ctx context.Context, queryName, query string, args ...any) error {
	_, err := s.tx.ExecContext(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("storage tx failed query %s: exec err: %w", queryName, err)
		return err
	}
	return nil
}

func (s *pgTx) ExecWithResult(ctx context.Context, queryName, query string, args ...any) (any, error) {
	result, err := s.tx.ExecContext(ctx, query, args)
	if err != nil {
		err = fmt.Errorf("storage tx failed query %s: exec with result err: %w", queryName, err)
		return nil, err
	}
	return result, nil
}

func (s *pgTx) QueryOne(ctx context.Context, dest any, queryName, query string, args ...any) error {
	err := s.tx.GetContext(ctx, dest, query, args...)
	if err != nil {
		err = fmt.Errorf("storage tx failed query %s: query one err: %w", queryName, err)
		return err
	}
	return nil
}

func (s *pgTx) QueryAll(ctx context.Context, queryName, query string, args ...any) (*sqlx.Rows, error) {
	rows, err := s.tx.QueryxContext(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("storage tx failed query %s: query all err: %w", queryName, err)
		return nil, err
	}

	return rows, nil
}

func (s *pgTx) Commit(_ context.Context) error {
	err := s.tx.Commit()
	if err != nil {
		err = fmt.Errorf("commit tx %s err: %w", s.txName, err)
		return err
	}
	return nil
}

func (s *pgTx) Rollback(_ context.Context) error {
	err := s.tx.Rollback()
	if err != nil {
		err = fmt.Errorf("rollback tx %s err: %w", s.txName, err)
		return err
	}

	return nil
}

func (s *pgTx) CommitRollback(_ context.Context, err error) error {
	if err != nil {
		if rollbackErr := s.tx.Rollback(); rollbackErr != nil {
			rollbackErr = fmt.Errorf("rollback CommitRollback tx %s err: %w", s.txName, rollbackErr)
			return rollbackErr
		}
		return nil
	}

	if commitErr := s.tx.Commit(); commitErr != nil {
		commitErr = fmt.Errorf("commit CommitRollback tx %s err: %w", s.txName, commitErr)
	}

	return nil
}

package database

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

func New(ctx context.Context, connURL string) (*DB, error) {
	conn, err := pgx.Connect(ctx, connURL)
	if err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.conn.Ping(ctx)
}

func (db *DB) Close(ctx context.Context) error {
	return db.conn.Close(ctx)
}

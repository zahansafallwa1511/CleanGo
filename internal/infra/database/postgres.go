package database

import (
	"context"
	"database/sql"

	"cleanandclean/internal/core/config"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db  *sql.DB
	cfg *config.DatabaseConfig
}

func NewPostgresDB(cfg *config.DatabaseConfig) *PostgresDB {
	return &PostgresDB{
		cfg: cfg,
	}
}

func (p *PostgresDB) Connect(ctx context.Context) error {
	db, err := sql.Open("postgres", p.cfg.URL)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(p.cfg.MaxOpenConns)
	db.SetMaxIdleConns(p.cfg.MaxIdleConns)
	db.SetConnMaxLifetime(p.cfg.ConnMaxLifetime)

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *PostgresDB) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

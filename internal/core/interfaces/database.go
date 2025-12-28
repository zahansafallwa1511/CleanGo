package interfaces

import (
	"context"
	"database/sql"
)

type IDatabase interface {
	Connect(ctx context.Context) error
	Close() error
	GetDB() *sql.DB
}

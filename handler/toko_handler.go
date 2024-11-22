package handler

import(
	"SecretCare/entity"
	"context"
	"fmt"

    _ "github.com/lib/pq" // PostgreSQL driver
	"database/sql"
)

type HandlerToko interface {
	CreateToko(ctx context.Context, toko entity.Toko) (int, error)
}

type handlerToko struct {
	ctx context.Context
	db  *sql.DB
}

// NewHandlerAuth membuat instance baru dari HandlerAuth
func NewHandlerToko(ctx context.Context, db *sql.DB) *handlerToko {
	return &handlerToko{
		ctx: ctx,
		db:  db,
	}
}

func (h *handlerToko) CreateToko(ctx context.Context, toko entity.Toko) (int, error) {
	// Insert into the database
	query := "INSERT INTO toko (nama) VALUES ($1) RETURNING id" // PostgreSQL uses $1 as a placeholder
	var lastInsertID int64
	err := h.db.QueryRowContext(ctx, query, toko.Nama).Scan(&lastInsertID)
	if err != nil {
		return 0, fmt.Errorf("error executing query: %v", err)
	}

	fmt.Printf("Toko baru berhasil dibuat: %v\n", toko.Nama)
	return int(lastInsertID), nil
}
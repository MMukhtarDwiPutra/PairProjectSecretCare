package handler

import(
	"SecretCare/entity"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type HandlerToko interface {
	CreateToko(ctx context.Context, toko entity.Toko) int64
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

func (h *handlerToko) CreateToko(ctx context.Context, toko entity.Toko) int64 {
	// Insert into the database
	result, err := h.db.Exec("INSERT INTO toko (nama) VALUES (?)", toko.Nama)
	if err != nil {
		fmt.Println("Error executing query:", err)
		fmt.Println()
		return 0
	}

	// Get the last inserted ID (if needed)
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		return 0
	}
	fmt.Printf("Toko baru berhasil dibuat: %v\n", toko.Nama)
	return lastInsertID
}
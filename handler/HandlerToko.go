package handler

import(
	"SecretCare/entity"
	"SecretCare/helpers"
	"SecretCare/utils"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type HandlerToko interface {
	CreateToko(ctx context.Context, toko entity.Toko) int64
}

func (h *handler) CreateToko(ctx context.Context, toko entity.Toko) int64 {
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
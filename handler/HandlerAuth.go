package handler

import (
	"SecretCare/entity"
	"SecretCare/helpers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type HandlerAuth interface {
	RegisterUser(user entity.Users)
	CreateToko(toko entity.Toko) int64
}

func (h *handler) RegisterUser(user entity.Users) {
	// Hash the password
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Failed to hash password:", err)
		return
	}

	// Insert into the database
	_, err = h.db.Exec("INSERT INTO users (username, password, full_name, toko_id, role) VALUES (?, ?, ?, ?, ?)", user.Username, hash, user.FullName, user.TokoID, user.Role)
	if err != nil {
		fmt.Println("Error executing query:", err)
		fmt.Println()
		return
	}
}

func (h *handler) CreateToko(toko entity.Toko) int64 {
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

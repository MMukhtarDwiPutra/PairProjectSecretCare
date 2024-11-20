package handler

import (
	"SecretCare/entity"
	"SecretCare/helpers"
	"SecretCare/utils"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type HandlerAuth interface {
	RegisterUser(ctx context.Context, user entity.Users)
	CreateToko(ctx context.Context, toko entity.Toko) int64
	Login(username, password string) (bool, string, context.Context)
}

func (h *handler) Login(username, password string) (bool, string, context.Context) {
	user, err := h.GetUserByUsername(username)
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return false, "", h.ctx
	}

	successLogin := helpers.CheckPasswordHash(password, user.Password)

	if successLogin {
		user := &entity.Users{ID: user.ID, Username: user.Username, FullName: user.FullName}
		h.ctx = utils.SetUserInContext(h.ctx, user) // Set user in the context
	}

	fmt.Println("")
	return successLogin, user.Role, h.ctx
}

func (h *handler) RegisterUser(ctx context.Context, user entity.Users) {
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

package handler

import (
	"SecretCare/entity"
	"SecretCare/helpers"
	"SecretCare/utils"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type HandlerAuth interface {
	RegisterUser(ctx context.Context, user entity.Users)
	Login(username, password string) (bool, string, context.Context, error)
}

type handlerAuth struct {
	handlerUser HandlerUser
	ctx context.Context
	db  *sql.DB
}

// NewHandlerAuth membuat instance baru dari HandlerAuth
func NewHandlerAuth(ctx context.Context, db *sql.DB) HandlerAuth {
	return &handlerAuth{
		handlerUser:    NewHandlerUser(ctx, db),
		ctx: ctx,
		db:  db,
	}
}

func (h *handlerAuth) Login(username, password string) (bool, string, context.Context, error) {
	user, err := h.handlerUser.GetUserByUsername(username)
	if err != nil {
		return false, "", h.ctx, err
	}

	successLogin := helpers.CheckPasswordHash(password, user.Password)

	if successLogin {
		user := &entity.Users{ID: user.ID, Username: user.Username, FullName: user.FullName, TokoID: user.TokoID}
		h.ctx = utils.SetUserInContext(h.ctx, user) // Set user in the context
	}

	return successLogin, user.Role, h.ctx, nil
}

func (h *handlerAuth) RegisterUser(ctx context.Context, user entity.Users) {
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
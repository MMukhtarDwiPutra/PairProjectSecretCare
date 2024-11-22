package handler

import (
	"SecretCare/entity"
	"SecretCare/helpers"
	"SecretCare/utils"
	"context"
	"fmt"

    _ "github.com/lib/pq" // PostgreSQL driver
	"database/sql"
)

type HandlerAuth interface {
	RegisterUser(ctx context.Context, user entity.Users) error 
	Login(username, password string) (bool, string, context.Context, error)
}

type handlerAuth struct {
	handlerUser HandlerUser
	ctx context.Context
	db  *sql.DB
}

// NewHandlerAuth membuat instance baru dari HandlerAuth
func NewHandlerAuth(ctx context.Context, db *sql.DB) *handlerAuth {
	return &handlerAuth{
		handlerUser:    NewHandlerUser(ctx, db),
		ctx: ctx,
		db:  db,
	}
}

func (h *handlerAuth) RegisterUser(ctx context.Context, user entity.Users) error {
	// Hash the password
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Insert into the database
	query := `
		INSERT INTO users (username, password, full_name, toko_id, role)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = h.db.ExecContext(ctx, query, user.Username, hash, user.FullName, user.TokoID, user.Role)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	fmt.Println("User registered successfully!")
	return nil
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

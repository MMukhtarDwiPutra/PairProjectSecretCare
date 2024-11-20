package handler

import (
	"SecretCare/entity"
	"SecretCare/utils"
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type HandlerUser interface {
	GetUserByUsername(username string) (*entity.Users, error)
	DeleteMyAccount(userId int) error
	UpdateMyAccount(username, password, fullName string) error
}

func (h *handler) GetUserByUsername(username string) (*entity.Users, error) {
	var user entity.Users
	row := h.db.QueryRow("SELECT id, username, full_name, role, password FROM users WHERE username = ?", username)

	if err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Role, &user.Password); err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}

	return &user, nil
}

func (h *handler) UpdateMyAccount(username, password, fullName *string) (context.Context, error) {
	user, ok := utils.GetUserFromContext(h.ctx)
	if !ok {
		return h.ctx, fmt.Errorf("user not found in context")
	}

	query := "UPDATE users SET "
	params := []interface{}{}

	if username != nil {
		query += "username = ?, "
		params = append(params, *username)
	}
	if password != nil {
		query += "password = ?, "
		params = append(params, *password)
	}
	if fullName != nil {
		query += "full_name = ?, "
		params = append(params, *fullName)
	}

	query = query[:len(query)-2]
	query += " WHERE id = ?"
	params = append(params, user.ID)

	_, err := h.db.Exec(query, params...)

	newUpdatedUser := &entity.Users{ID: user.ID}
	if username != nil {
		newUpdatedUser.Username = *username
	}
	if fullName != nil {
		newUpdatedUser.FullName = *fullName
	}

	h.ctx = utils.SetUserInContext(h.ctx, newUpdatedUser)
	if err != nil {
		return h.ctx, fmt.Errorf("failed to update user: %w", err)
	}
	return h.ctx, nil
}

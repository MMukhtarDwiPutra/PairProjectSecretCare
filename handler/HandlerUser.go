package handler

import (
	"SecretCare/entity"
	"SecretCare/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type HandlerUser interface {
	GetUserByUsername(username string) (*entity.Users, error)
}

func (h *handler) GetUserByUsername(username string) (*entity.Users, error) {
	var user entity.Users
	fmt.Print(h.ctx)
	users, ok := utils.GetUserFromContext(h.ctx)
	if ok {
		fmt.Print(users.FullName, users.ID, "test 123213")
	}

	row := h.db.QueryRow("SELECT id, username, full_name, role, password FROM users WHERE username = ?", username)

	if err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Role, &user.Password); err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}

	return &user, nil
}

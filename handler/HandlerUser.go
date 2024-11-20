package handler

import(
	_ "github.com/go-sql-driver/mysql"
	"SecretCare/entity"
)

type HandlerUser interface{
	GetUserByUsername(username string) entity.Users
}

func (h *handler) GetUserByUsername(username string) entity.Users{
	var user entity.Users

	row := h.db.QueryRow("SELECT role, password FROM users WHERE username = ?", username)

	row.Scan(&user.Role, &user.Password)

	return user
}

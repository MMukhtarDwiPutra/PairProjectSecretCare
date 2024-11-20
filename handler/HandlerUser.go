package handler

import(
	// "ProjectDatabase/entity"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"SecretCare/helpers"
	"SecretCare/entity"
)

type HandlerUser interface{
	GetUserByUsername(username string) entity.Users
	RegisterUser(user entity.Users)
	CreateToko(toko entity.Toko) int64
}

type handlerUser struct{
	db *sql.DB
}

func NewHandler(db *sql.DB) *handlerUser{
	return &handlerUser{db}
}

func (h *handlerUser) GetUserByUsername(username string) entity.Users{
	var user entity.Users

	row := h.db.QueryRow("SELECT role, password FROM users WHERE username = ?", username)

	row.Scan(&user.Role, &user.Password)

	return user
}

func (h *handlerUser) RegisterUser(user entity.Users){
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

func (h *handlerUser) CreateToko(toko entity.Toko) int64{
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
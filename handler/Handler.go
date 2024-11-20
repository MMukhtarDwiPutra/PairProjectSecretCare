package handler

import(
	// "ProjectDatabase/entity"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
	// "SecretCare/helpers"
	// "SecretCare/entity"
)

type Handler interface{
}

type handler struct{
	db *sql.DB
}

func NewHandler(db *sql.DB) *handler{
	return &handler{db}
}
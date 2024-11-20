package handler

import "database/sql"

type Handler struct {
	Auth HandlerAuth
	User HandlerUser
}

type handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) Handler {
	return Handler{
		Auth: &handler{db: db},
		User: &handler{db: db},
	}
}

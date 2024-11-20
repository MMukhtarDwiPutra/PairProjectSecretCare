package handler

import (
	"context"
	"database/sql"
)

// Define a single handler struct that handles both Auth and User logic
type Handler struct {
	Handler *handler
}

// Define the handler struct which holds context and DB connection
type handler struct {
	ctx context.Context
	db  *sql.DB
}

// NewHandler creates a new Handler with both Auth and User logic combined.
func NewHandler(ctx context.Context, db *sql.DB) Handler {
	return Handler{
		Handler: &handler{ctx: ctx, db: db},
	}
}

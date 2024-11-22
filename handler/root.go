package handler

import (
	"context"
	"database/sql"
)

// Handler adalah root struct yang menggabungkan semua sub-handler
type Handler struct {
	Auth    HandlerAuth
	Toko    HandlerToko
	User    HandlerUser
	Product HandlerProduct
	Cart    HandlerCart
	Order   HandlerOrder
	ctx context.Context
	db 	*sql.DB
}

// NewHandler membuat instance dari root Handler dan semua sub-handler
func NewHandler(ctx context.Context, db *sql.DB) *Handler {
	return &Handler{
		Auth:    NewHandlerAuth(ctx, db),
		Toko:    NewHandlerToko(ctx, db),
		User:    NewHandlerUser(ctx, db),
		Product: NewHandlerProduct(ctx, db),
		Cart:    NewHandlerCart(db),
		Order:   NewHandlerOrder(db), 
		ctx: ctx,
		db: db,
	}
}

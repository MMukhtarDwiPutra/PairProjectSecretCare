package handler

import (
	"database/sql"
	"fmt"
	"context"
)

type HandlerOrder interface {
	CreateNewOrder(cartID int) error
	UpdateCartStatus(cartID int, status string) error
	Checkout(userID int) error
}

type handlerOrder struct {
	db  *sql.DB
	ctx context.Context
}

func NewHandlerOrder(ctx context.Context, db *sql.DB) HandlerOrder {
	return &handlerOrder{db: db, ctx: ctx}
}

func (h *handlerOrder) CreateNewOrder(cartID int) error {
	// Insert a new order into the database
	query := `
		INSERT INTO orders (status, order_date, cart_id)
		VALUES ('Shipped', NOW(), ?)
	`
	_, err := h.db.ExecContext(h.ctx, query, cartID)
	if err != nil {
		return fmt.Errorf("failed to create new order: %v", err)
	}

	return nil
}

func (h *handlerOrder) UpdateCartStatus(userID int, status string) error {
	// Update the status of the cart
	query := `
		UPDATE carts
		SET status = ?
		WHERE user_id = ?
	`
	_, err := h.db.ExecContext(h.ctx, query, status, userID)
	if err != nil {
		return fmt.Errorf("failed to update cart status: %v", err)
	}

	return nil
}

func (h *handlerOrder) Checkout(userID int) error {
	// Get the active cart for the user
	var cartID int
	query := `
		SELECT id
		FROM carts
		WHERE user_id = ? AND status = 'Active'
	`
	err := h.db.QueryRowContext(h.ctx, query, userID).Scan(&cartID)
	if err != nil {
		fmt.Println("the error from querying row context")

		if err == sql.ErrNoRows {
			return fmt.Errorf("no active cart found for user ID %d", userID)
		}
		return fmt.Errorf("failed to retrieve active cart: %v", err)
	}

	// Create a new order
	err = h.CreateNewOrder(cartID)
	if err != nil {
		fmt.Printf("Error creating order: %v\n", err) // Add logging
		return fmt.Errorf("failed to create order: %v", err)
	}

	// Update the cart status to 'Checked_Out'
	err = h.UpdateCartStatus(userID, "Checked Out")
	if err != nil {
		fmt.Println("the error from update cart status");
		return fmt.Errorf("failed to update cart status: %v", err)
	}

	return nil
}
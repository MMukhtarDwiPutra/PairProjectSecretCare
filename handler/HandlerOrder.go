package handler

import (
	"database/sql"
	"fmt"
)

type HandlerOrder interface {
	CreateNewOrder(cartID int) error
	UpdateCartStatus(cartID int, status string) error
	Checkout(userID int) error
}

type handlerOrder struct {
	db *sql.DB
}

func NewHandlerOrder(db *sql.DB) HandlerOrder {
	return &handlerOrder{db: db}
}

func (h *handlerOrder) CreateNewOrder(cartID int) error {
	// Insert a new order into the database
	query := `
		INSERT INTO orders (status, order_date, cart_id)
		VALUES ('Sudah Dikirim', NOW(), ?)
	`
	_, err := h.db.Exec(query, cartID)
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
	_, err := h.db.Exec(query, status, userID)
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
	err := h.db.QueryRow(query, userID).Scan(&cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no active cart found for user ID %d", userID)
		}
		return fmt.Errorf("failed to retrieve active cart: %v", err)
	}

	// Create a new order
	err = h.CreateNewOrder(cartID)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	// Update the cart status to 'Checked_Out'
	err = h.UpdateCartStatus(userID, "Checked Out")
	if err != nil {
		return fmt.Errorf("failed to update cart status: %v", err)
	}

	return nil
}
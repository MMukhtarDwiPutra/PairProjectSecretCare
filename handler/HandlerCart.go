package handler

import (
	"database/sql"
	"fmt"
)

type HandlerCart interface {
	AddCart(userID int, productID int, qty int, priceAtPurchase float64) error
}

type handlerCart struct {
	db *sql.DB
}

// NewHandlerCart creates a new instance of HandlerCart
func NewHandlerCart(db *sql.DB) HandlerCart {
	return &handlerCart{db: db}
}

func (h *handlerCart) AddCart(userID int, productID int, qty int, priceAtPurchase float64) error {
	// Check if there's an active cart for the user
	var cartID int
	query := "SELECT id FROM carts WHERE user_id = ? AND status = 'Active'"
	err := h.db.QueryRow(query, userID).Scan(&cartID)

	if err != nil {
		if err == sql.ErrNoRows {
			// No active cart, create a new one
			cartID, err = h.CreateNewCart(userID)
			if err != nil {
				return fmt.Errorf("failed to create new cart: %v", err)
			}
		} else {
			// Some other database error
			return fmt.Errorf("error checking active cart: %v", err)
		}
	}

	// Add item to cart
	err = h.CreateNewCartItems(cartID, productID, qty, priceAtPurchase)
	if err != nil {
		return fmt.Errorf("failed to create new cart items: %v", err)
	}

	fmt.Println("Cart item added successfully")
	return nil
}

// CreateNewCart creates a new cart for the user and returns the cart ID
func (h *handlerCart) CreateNewCart(userID int) (int, error) {
	result, err := h.db.Exec("INSERT INTO carts (status, user_id) VALUES ('Active', ?)", userID)
	if err != nil {
		return 0, fmt.Errorf("failed to create cart: %v", err)
	}

	cartID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %v", err)
	}

	return int(cartID), nil
}

// CreateNewCartItems adds a new item to the cart
func (h *handlerCart) CreateNewCartItems(cartID, productID, qty int, priceAtPurchase float64) error {
	_, err := h.db.Exec("INSERT INTO cart_items (cart_id, product_id, qty, price_at_purchase) VALUES (?, ?, ?, ?)",
		cartID, productID, qty, priceAtPurchase)
	if err != nil {
		return fmt.Errorf("failed to add item to cart: %v", err)
	}
	return nil
}
package handler

import (
	"context"
	"database/sql"
	"fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
)

type HandlerCart interface {
	AddCart(userID int, productID int, qty int, priceAtPurchase float64) error
	ShowCart(userID int) ([]struct {
		ProductName string
		Quantity    int
		Status      string
	}, error)
	DeleteAllCartItemsActive(userID int) error
	DeleteCartItemByID(cartItemID int) error
	GetActiveCartItems(userID int) ([]struct {
		ID         int
		ProductName string
		Quantity   int
	}, error)
	UpdateQuantityCart(cartItemID int, newQuantity int) error
}

type handlerCart struct {
	db  *sql.DB
	ctx context.Context
}

func NewHandlerCart(ctx context.Context, db *sql.DB) HandlerCart {
	return &handlerCart{db: db, ctx: ctx}
}

func (h *handlerCart) AddCart(userID int, productID int, qty int, priceAtPurchase float64) error {
	var cartID int
	query := "SELECT id FROM carts WHERE user_id = $1 AND status = 'Active'"
	err := h.db.QueryRowContext(h.ctx, query, userID).Scan(&cartID)

	if err != nil {
		if err == sql.ErrNoRows {
			cartID, err = h.CreateNewCart(userID)
			if err != nil {
				return fmt.Errorf("failed to create new cart: %w", err)
			}
		} else {
			return fmt.Errorf("error checking active cart: %w", err)
		}
	}

	err = h.CreateNewCartItems(cartID, productID, qty, priceAtPurchase)
	if err != nil {
		return fmt.Errorf("failed to create new cart items: %w", err)
	}

	fmt.Println("Cart item added successfully")
	return nil
}

func (h *handlerCart) CreateNewCart(userID int) (int, error) {
	query := "INSERT INTO carts (status, user_id) VALUES ('Active', $1) RETURNING id"
	var cartID int
	err := h.db.QueryRowContext(h.ctx, query, userID).Scan(&cartID)
	if err != nil {
		return 0, fmt.Errorf("failed to create cart: %w", err)
	}
	return cartID, nil
}

func (h *handlerCart) CreateNewCartItems(cartID, productID, qty int, priceAtPurchase float64) error {
	query := `
		INSERT INTO cart_items (cart_id, product_id, qty, price_at_purchase) 
		VALUES ($1, $2, $3, $4)
	`
	_, err := h.db.ExecContext(h.ctx, query, cartID, productID, qty, priceAtPurchase)
	if err != nil {
		return fmt.Errorf("failed to add item to cart: %w", err)
	}
	return nil
}

func (h *handlerCart) ShowCart(userID int) ([]struct {
	ProductName string
	Quantity    int
	Status      string
}, error) {
	query := `
		SELECT p.nama AS product_name, ci.qty AS quantity, c.status AS status
		FROM carts c
		JOIN cart_items ci ON c.id = ci.cart_id
		JOIN products p ON ci.product_id = p.id
		WHERE c.user_id = $1
	`

	rows, err := h.db.QueryContext(h.ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart data: %w", err)
	}
	defer rows.Close()

	var cartItems []struct {
		ProductName string
		Quantity    int
		Status      string
	}

	for rows.Next() {
		var item struct {
			ProductName string
			Quantity    int
			Status      string
		}
		err := rows.Scan(&item.ProductName, &item.Quantity, &item.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		cartItems = append(cartItems, item)
	}

	return cartItems, nil
}

func (h *handlerCart) DeleteAllCartItemsActive(userID int) error {
	var cartID int
	query := `
		SELECT id
		FROM carts
		WHERE user_id = $1 AND status = 'Active'
	`
	err := h.db.QueryRowContext(h.ctx, query, userID).Scan(&cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no active cart found for user ID %d", userID)
		}
		return fmt.Errorf("failed to retrieve active cart: %w", err)
	}

	deleteQuery := `
		DELETE FROM cart_items
		WHERE cart_id = $1
	`
	_, err = h.db.ExecContext(h.ctx, deleteQuery, cartID)
	if err != nil {
		return fmt.Errorf("failed to delete cart items: %w", err)
	}

	return nil
}

func (h *handlerCart) DeleteCartItemByID(cartItemID int) error {
	query := `
		DELETE FROM cart_items
		WHERE id = $1
	`
	_, err := h.db.ExecContext(h.ctx, query, cartItemID)
	if err != nil {
		return fmt.Errorf("failed to delete cart item with ID %d: %w", cartItemID, err)
	}

	return nil
}

func (h *handlerCart) GetActiveCartItems(userID int) ([]struct {
	ID         int
	ProductName string
	Quantity   int
}, error) {
	query := `
		SELECT ci.id, p.nama AS product_name, ci.qty
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		JOIN carts c ON ci.cart_id = c.id
		WHERE c.user_id = $1 AND c.status = 'Active'
	`

	rows, err := h.db.QueryContext(h.ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %w", err)
	}
	defer rows.Close()

	var cartItems []struct {
		ID         int
		ProductName string
		Quantity   int
	}

	for rows.Next() {
		var item struct {
			ID         int
			ProductName string
			Quantity   int
		}
		err := rows.Scan(&item.ID, &item.ProductName, &item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		cartItems = append(cartItems, item)
	}

	return cartItems, nil
}

func (h *handlerCart) UpdateQuantityCart(cartItemID int, newQuantity int) error {
	if newQuantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	query := `
		UPDATE cart_items
		SET qty = $1
		WHERE id = $2
	`
	_, err := h.db.ExecContext(h.ctx, query, newQuantity, cartItemID)
	if err != nil {
		return fmt.Errorf("failed to update cart item quantity: %w", err)
	}

	return nil
}
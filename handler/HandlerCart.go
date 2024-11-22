package handler

import (
	"database/sql"
	"fmt"
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
		WHERE c.user_id = ?
	`

	rows, err := h.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart data: %v", err)
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
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		cartItems = append(cartItems, item)
	}

	return cartItems, nil
}

func (h *handlerCart) DeleteAllCartItemsActive(userID int) error {
	// Get the active cart ID for the user
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

	// Delete all items in the active cart
	deleteQuery := `
		DELETE FROM cart_items
		WHERE cart_id = ?
	`
	_, err = h.db.Exec(deleteQuery, cartID)
	if err != nil {
		return fmt.Errorf("failed to delete cart items: %v", err)
	}

	return nil
}

func (h *handlerCart) DeleteCartItemByID(cartItemID int) error {
	query := `
		DELETE FROM cart_items
		WHERE id = ?
	`
	_, err := h.db.Exec(query, cartItemID)
	if err != nil {
		return fmt.Errorf("failed to delete cart item with ID %d: %v", cartItemID, err)
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
		WHERE c.user_id = ? AND c.status = 'Active'
	`
	rows, err := h.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart items: %v", err)
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
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		cartItems = append(cartItems, item)
	}

	return cartItems, nil
}
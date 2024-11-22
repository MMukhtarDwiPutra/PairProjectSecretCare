package handler

import (
	"context"
	"database/sql"
	"fmt"
	"SecretCare/entity"
	
    _ "github.com/lib/pq" // PostgreSQL driver
)

type HandlerOrder interface {
	CreateNewOrder(cartID int) error
	UpdateCartStatus(userID int, status string) error
	Checkout(userID int) error
	GetAllOrderByTokoId(tokoID int) ([]entity.Order, error)
	UpdateStatusOrder(id int, status string) error
}

type handlerOrder struct {
	db  *sql.DB
	ctx context.Context
}

func NewHandlerOrder(ctx context.Context, db *sql.DB) HandlerOrder {
	return &handlerOrder{db: db, ctx: ctx}
}

func (h *handlerOrder) CreateNewOrder(cartID int) error {
	query := `
		INSERT INTO orders (status, order_date, cart_id)
		VALUES ('Shipped', NOW(), $1)
	`
	_, err := h.db.ExecContext(h.ctx, query, cartID)
	if err != nil {
		return fmt.Errorf("failed to create new order: %w", err)
	}
	return nil
}

func (h *handlerOrder) UpdateCartStatus(userID int, status string) error {
	query := `
		UPDATE carts
		SET status = $1
		WHERE user_id = $2
	`
	_, err := h.db.ExecContext(h.ctx, query, status, userID)
	if err != nil {
		return fmt.Errorf("failed to update cart status: %w", err)
	}
	return nil
}

func (h *handlerOrder) Checkout(userID int) error {
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

	err = h.CreateNewOrder(cartID)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	err = h.UpdateCartStatus(userID, "Checked Out")
	if err != nil {
		return fmt.Errorf("failed to update cart status: %w", err)
	}

	return nil
}

func (h *handlerOrder) GetAllOrderByTokoId(tokoID int) ([]entity.Order, error) {
	var orders []entity.Order
	query := `
		SELECT DISTINCT 
			orders.id, 
			products.nama, 
			orders.status, 
			orders.order_date, 
			orders.cart_id
		FROM orders
		JOIN carts 
			ON carts.id = orders.cart_id
		JOIN cart_items 
			ON cart_items.cart_id = carts.id
		JOIN products 
			ON cart_items.product_id = products.id
		WHERE products.toko_id = $1
	`
	rows, err := h.db.QueryContext(h.ctx, query, tokoID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve orders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.NamaProduct, &order.Status, &order.OrderDate, &order.CartID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return orders, nil
}

func (h *handlerOrder) UpdateStatusOrder(id int, status string) error {
	query := `
		UPDATE orders
		SET status = $1
		WHERE id = $2
	`
	result, err := h.db.ExecContext(h.ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update orders: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated for orders.id: %d", id)
	}

	return nil
}
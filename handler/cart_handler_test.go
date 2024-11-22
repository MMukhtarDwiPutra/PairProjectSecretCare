package handler

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddCart_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	// Mock queries
	mock.ExpectQuery("SELECT id FROM carts WHERE user_id = \\? AND status = 'Active'").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	mock.ExpectExec("INSERT INTO carts \\(status, user_id\\) VALUES \\('Active', \\?\\)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO cart_items \\(cart_id, product_id, qty, price_at_purchase\\) VALUES \\(\\?, \\?, \\?, \\?\\)").
		WithArgs(1, 101, 2, 50.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = handler.AddCart(1, 101, 2, 50.0)
	assert.Nil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddCart_NoActiveCartError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	// Mock queries
	mock.ExpectQuery("SELECT id FROM carts WHERE user_id = \\? AND status = 'Active'").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectExec("INSERT INTO carts \\(status, user_id\\) VALUES \\('Active', \\?\\)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO cart_items \\(cart_id, product_id, qty, price_at_purchase\\) VALUES \\(\\?, \\?, \\?, \\?\\)").
		WithArgs(1, 101, 2, 50.0).
		WillReturnError(fmt.Errorf("failed to add cart item"))

	err = handler.AddCart(1, 101, 2, 50.0)
	assert.NotNil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestShowCart_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	mock.ExpectQuery("SELECT p.nama AS product_name, ci.qty AS quantity, c.status AS status FROM carts c JOIN cart_items ci ON c.id = ci.cart_id JOIN products p ON ci.product_id = p.id WHERE c.user_id = \\?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"product_name", "quantity", "status"}).
			AddRow("Keyboard", 2, "Active").
			AddRow("Mouse", 3, "Active"))

	result, err := handler.ShowCart(1)
	assert.Nil(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Keyboard", result[0].ProductName)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAllCartItemsActive_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	mock.ExpectQuery("SELECT id FROM carts WHERE user_id = \\? AND status = 'Active'").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectExec("DELETE FROM cart_items WHERE cart_id = \\?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 2))

	err = handler.DeleteAllCartItemsActive(1)
	assert.Nil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCartItemByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	mock.ExpectExec("DELETE FROM cart_items WHERE id = \\?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = handler.DeleteCartItemByID(1)
	assert.Nil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActiveCartItems_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	mock.ExpectQuery("SELECT ci.id, p.nama AS product_name, ci.qty FROM cart_items ci JOIN products p ON ci.product_id = p.id JOIN carts c ON ci.cart_id = c.id WHERE c.user_id = \\? AND c.status = 'Active'").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "product_name", "qty"}).
			AddRow(1, "Keyboard", 2).
			AddRow(2, "Mouse", 3))

	result, err := handler.GetActiveCartItems(1)
	assert.Nil(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Keyboard", result[0].ProductName)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateQuantityCart_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	mock.ExpectExec("UPDATE cart_items SET qty = \\? WHERE id = \\?").
		WithArgs(3, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = handler.UpdateQuantityCart(1, 3)
	assert.Nil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateQuantityCart_InvalidQuantity(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerCart(ctx, db)

	err = handler.UpdateQuantityCart(1, -1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "quantity must be greater than zero")
}

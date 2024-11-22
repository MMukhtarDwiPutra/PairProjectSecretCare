package handler

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewOrder_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock ExecContext for order creation
	mock.ExpectExec(`INSERT INTO orders \(status, order_date, cart_id\) VALUES \('Shipped', NOW\(\), \?\)`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = handler.CreateNewOrder(1)
	assert.Nil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateNewOrder_Failure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock ExecContext to simulate failure
	mock.ExpectExec(`INSERT INTO orders \(status, order_date, cart_id\) VALUES \('Shipped', NOW\(\), \?\)`).
		WithArgs(1).
		WillReturnError(fmt.Errorf("database error"))

	err = handler.CreateNewOrder(1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to create new order")

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCartStatus_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock ExecContext for cart status update
	mock.ExpectExec(`UPDATE carts SET status = \? WHERE user_id = \?`).
		WithArgs("Checked Out", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = handler.UpdateCartStatus(1, "Checked Out")
	assert.Nil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCartStatus_Failure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock ExecContext to simulate failure
	mock.ExpectExec(`UPDATE carts SET status = \? WHERE user_id = \?`).
		WithArgs("Checked Out", 1).
		WillReturnError(fmt.Errorf("update failed"))

	err = handler.UpdateCartStatus(1, "Checked Out")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to update cart status")

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckout_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock QueryRowContext for active cart retrieval
	mock.ExpectQuery(`SELECT id FROM carts WHERE user_id = \? AND status = 'Active'`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Mock ExecContext for order creation
	mock.ExpectExec(`INSERT INTO orders \(status, order_date, cart_id\) VALUES \('Shipped', NOW\(\), \?\)`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock ExecContext for cart status update
	mock.ExpectExec(`UPDATE carts SET status = \? WHERE user_id = \?`).
		WithArgs("Checked Out", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = handler.Checkout(1)
	assert.Nil(t, err)

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckout_NoActiveCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock QueryRowContext to simulate no active cart
	mock.ExpectQuery(`SELECT id FROM carts WHERE user_id = \? AND status = 'Active'`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	err = handler.Checkout(1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no active cart found")

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckout_CreateOrderFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock QueryRowContext for active cart retrieval
	mock.ExpectQuery(`SELECT id FROM carts WHERE user_id = \? AND status = 'Active'`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Mock ExecContext to simulate order creation failure
	mock.ExpectExec(`INSERT INTO orders \(status, order_date, cart_id\) VALUES \('Shipped', NOW\(\), \?\)`).
		WithArgs(1).
		WillReturnError(fmt.Errorf("order creation failed"))

	err = handler.Checkout(1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to create order")

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckout_UpdateCartStatusFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	handler := NewHandlerOrder(ctx, db)

	// Mock QueryRowContext for active cart retrieval
	mock.ExpectQuery(`SELECT id FROM carts WHERE user_id = \? AND status = 'Active'`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Mock ExecContext for order creation
	mock.ExpectExec(`INSERT INTO orders \(status, order_date, cart_id\) VALUES \('Shipped', NOW\(\), \?\)`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock ExecContext to simulate cart status update failure
	mock.ExpectExec(`UPDATE carts SET status = \? WHERE user_id = \?`).
		WithArgs("Checked Out", 1).
		WillReturnError(fmt.Errorf("cart update failed"))

	err = handler.Checkout(1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to update cart status")

	// Verify expectations
	assert.NoError(t, mock.ExpectationsWereMet())
}
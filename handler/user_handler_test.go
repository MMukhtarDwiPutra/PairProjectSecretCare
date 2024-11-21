package handler_test

import (
	"SecretCare/entity"
	"SecretCare/handler"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestHandlerUser_GetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	h := handler.NewHandlerUser(context.Background(), db)

	mock.ExpectQuery("SELECT id, username, full_name, role, password, toko_id FROM users WHERE username = ?").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "full_name", "role", "password", "toko_id"}).
			AddRow(1, "testuser", "Test User", "User", "hashed_password", 1))

	user, err := h.GetUserByUsername("testuser")

	assert.NoError(t, err)
	assert.Equal(t, &entity.Users{
		ID:       1,
		Username: "testuser",
		FullName: "Test User",
		Role:     "User",
		Password: "hashed_password",
		TokoID:   1,
	}, user)

	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestUpdateMyAccount_UpdateUsernameOnly(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	h := handler.NewHandlerUser(ctx, db)

	newUsername := "new_username"

	mock.ExpectExec(`UPDATE users SET username = \? WHERE id = \?`).
		WithArgs(newUsername, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = h.UpdateMyAccount(1, &newUsername, nil, nil)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateMyAccount_UpdateUsernameAndPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	h := handler.NewHandlerUser(ctx, db)

	newUsername := "new_username"
	newPassword := "new_password"

	mock.ExpectExec(`UPDATE users SET username = \?, password = \? WHERE id = \?`).
		WithArgs(newUsername, newPassword, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = h.UpdateMyAccount(1, &newUsername, &newPassword, nil)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateMyAccount_UpdateAllFields(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	h := handler.NewHandlerUser(ctx, db)

	newUsername := "new_username"
	newPassword := "new_password"
	newFullName := "new_full_name"

	mock.ExpectExec(`UPDATE users SET username = \?, password = \?, full_name = \? WHERE id = \?`).
		WithArgs(newUsername, newPassword, newFullName, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = h.UpdateMyAccount(1, &newUsername, &newPassword, &newFullName)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateMyAccount_NoFieldsToUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	h := handler.NewHandlerUser(ctx, db)

	newUsername := "new_username"

	mock.ExpectExec(`UPDATE users SET username = \? WHERE id = \?`).
		WithArgs(newUsername, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = h.UpdateMyAccount(1, &newUsername, nil, nil)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHandlerUser_DeleteMyAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	h := handler.NewHandlerUser(context.Background(), db)

	mock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = h.DeleteMyAccount(1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReportBuyerSpending(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	h := handler.NewHandlerUser(ctx, db)

	rows := sqlmock.NewRows([]string{"order_id", "user_id", "full_name", "total_spending", "total_qty"}).
		AddRow(1, 101, "John Doe", 200.0, 2).
		AddRow(2, 101, "John Doe", 300.0, 3)

	mock.ExpectQuery(`SELECT o.id AS order_id, u.id AS user_id, u.full_name, SUM\(ci.qty \* ci.price_at_purchase\) AS total_spending, SUM\(ci.qty\) AS total_qty`).
		WithArgs(101).
		WillReturnRows(rows)

	result, err := h.ReportBuyerSpending(101)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, entity.UserBuyerReport{OrderID: 1, UserID: 101, FullName: "John Doe", TotalSpending: 200.0, TotalQuantity: 2}, result[0])
	assert.Equal(t, entity.UserBuyerReport{OrderID: 2, UserID: 101, FullName: "John Doe", TotalSpending: 300.0, TotalQuantity: 3}, result[1])

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReportUserWithHighestSpending(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	h := handler.NewHandlerUser(ctx, db)

	rows := sqlmock.NewRows([]string{"user_id", "full_name", "total_spending"}).
		AddRow(201, "Alice", 500.0).
		AddRow(202, "Bob", 400.0)

	mock.ExpectQuery(`SELECT u.id AS user_id, u.full_name, SUM\(ci.qty \* ci.price_at_purchase\) AS total_spending`).
		WithArgs(1).
		WillReturnRows(rows)

	result, err := h.ReportUserWithHighestSpending(1)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, entity.UserReportHighestSpending{UserId: 201, FullName: "Alice", TotalSpending: 500.0}, result[0])
	assert.Equal(t, entity.UserReportHighestSpending{UserId: 202, FullName: "Bob", TotalSpending: 400.0}, result[1])

	assert.NoError(t, mock.ExpectationsWereMet())
}

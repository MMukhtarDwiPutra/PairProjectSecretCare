package handler

import (
	"SecretCare/entity"
	"context"
	"fmt"
	"log"
	"database/sql"
    _ "github.com/lib/pq" // PostgreSQL driver
)

type HandlerUser interface {
	GetUserByUsername(username string) (*entity.Users, error)
	DeleteMyAccount(userId int) error
	UpdateMyAccount(userId int, username, password, fullName *string) error
	ReportBuyerSpending(userId int) ([]entity.UserBuyerReport, error)
	ReportUserWithHighestSpending(tokoId int) ([]entity.UserReportHighestSpending, error)
}

type handlerUser struct {
	ctx context.Context
	db  *sql.DB
}

// NewHandlerUser membuat instance baru dari HandlerUser
func NewHandlerUser(ctx context.Context, db *sql.DB) *handlerUser {
	return &handlerUser{
		ctx: ctx,
		db:  db,
	}
}

func (h *handlerUser) GetUserByUsername(username string) (*entity.Users, error) {
	var user entity.Users
	query := `
		SELECT id, username, full_name, role, password, toko_id 
		FROM users 
		WHERE username = $1
	`
	row := h.db.QueryRowContext(h.ctx, query, username)

	if err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Role, &user.Password, &user.TokoID); err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}

	return &user, nil
}

func (h *handlerUser) UpdateMyAccount(userId int, username, password, fullName *string) error {
	query := "UPDATE users SET "
	params := []interface{}{}

	if username != nil {
		query += "username = $1, "
		params = append(params, *username)
	}
	if password != nil {
		query += "password = $2, "
		params = append(params, *password)
	}
	if fullName != nil {
		query += "full_name = $3, "
		params = append(params, *fullName)
	}

	// Remove trailing comma and space
	query = query[:len(query)-2]
	query += " WHERE id = $4"
	params = append(params, userId)

	_, err := h.db.ExecContext(h.ctx, query, params...)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (h *handlerUser) DeleteMyAccount(userId int) error {
	query := `
		DELETE FROM users 
		WHERE id = $1
	`
	_, err := h.db.ExecContext(h.ctx, query, userId)
	if err != nil {
		return fmt.Errorf("failed to delete account: %w", err)
	}

	return nil
}

func (h *handlerUser) ReportBuyerSpending(userId int) ([]entity.UserBuyerReport, error) {
	var orders []entity.UserBuyerReport
	query := `
		SELECT 
			o.id AS order_id, 
			u.id AS user_id, 
			u.full_name, 
			COALESCE(SUM(ci.qty * ci.price_at_purchase), 0) AS total_spending, 
			COALESCE(SUM(ci.qty), 0) AS total_qty
		FROM 
			users u
		JOIN 
			carts c ON u.id = c.user_id
		JOIN 
			cart_items ci ON c.id = ci.cart_id 
		JOIN 
			orders o ON c.id = o.cart_id
		WHERE 
			u.role = 'Pembeli'
			AND c.status = 'Checked Out'
			AND u.id = $1
		GROUP BY 
			o.id, u.id
		ORDER BY 
			total_spending DESC
	`
	rows, err := h.db.QueryContext(h.ctx, query, userId)
	if err != nil {
		log.Printf("Error fetching buyer spending records for user ID %d: %v", userId, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderReport entity.UserBuyerReport
		err = rows.Scan(&orderReport.OrderID, &orderReport.UserID, &orderReport.FullName, &orderReport.TotalSpending, &orderReport.TotalQuantity)
		if err != nil {
			log.Printf("Error scanning buyer spending record for user ID %d: %v", userId, err)
			return nil, err
		}
		orders = append(orders, orderReport)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating buyer spending records for user ID %d: %v", userId, err)
		return nil, err
	}

	return orders, nil
}

func (h *handlerUser) ReportUserWithHighestSpending(tokoId int) ([]entity.UserReportHighestSpending, error) {
	var users []entity.UserReportHighestSpending
	query := `
		SELECT 
			u.id AS user_id, 
			u.full_name, 
			COALESCE(SUM(ci.qty * ci.price_at_purchase), 0) AS total_spending
		FROM 
			users u
		JOIN 
			carts c ON u.id = c.user_id
		JOIN 
			cart_items ci ON c.id = ci.cart_id
		WHERE 
			c.status = 'Checked Out'
			AND u.toko_id = $1
		GROUP BY 
			u.id
		ORDER BY 
			total_spending DESC
		LIMIT 10
	`

	rows, err := h.db.QueryContext(h.ctx, query, tokoId)
	if err != nil {
		log.Printf("Error fetching records: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.UserReportHighestSpending
		err = rows.Scan(&user.UserId, &user.FullName, &user.TotalSpending)
		if err != nil {
			log.Printf("Error scanning record: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, err
	}

	return users, nil
}

package handler

import (
	"SecretCare/entity"
	"SecretCare/utils"
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
type HandlerUser interface {
	GetUserByUsername(username string) (*entity.Users, error)
	DeleteMyAccount(userId int) error
	UpdateMyAccount(username, password, fullName string) error
	ReportBuyerSpending() error
	ReportSellerSpending() error
}

func (h *handler) GetUserByUsername(username string) (*entity.Users, error) {
	var user entity.Users
	row := h.db.QueryRow("SELECT id, username, full_name, role, password, toko_id FROM users WHERE username = ?", username)
  
	if err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Role, &user.Password, &user.TokoID); err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}


	return &user, nil
}

func (h *handler) UpdateMyAccount(username, password, fullName *string) (context.Context, error) {
	user, ok := utils.GetUserFromContext(h.ctx)
	if !ok {
		return h.ctx, fmt.Errorf("user not found in context")
	}

	query := "UPDATE users SET "
	params := []interface{}{}

	if username != nil {
		query += "username = ?, "
		params = append(params, *username)
	}
	if password != nil {
		query += "password = ?, "
		params = append(params, *password)
	}
	if fullName != nil {
		query += "full_name = ?, "
		params = append(params, *fullName)
	}

	query = query[:len(query)-2]
	query += " WHERE id = ?"
	params = append(params, user.ID)

	_, err := h.db.Exec(query, params...)

	newUpdatedUser := &entity.Users{ID: user.ID, TokoID: user.TokoID}
	if username != nil {
		newUpdatedUser.Username = *username
	}
	if fullName != nil {
		newUpdatedUser.FullName = *fullName
	}

	h.ctx = utils.SetUserInContext(h.ctx, newUpdatedUser)
	if err != nil {
		return h.ctx, fmt.Errorf("failed to update user: %w", err)
	}
	return h.ctx, nil
}

func (h *handler) DeleteMyAccount() (context.Context, error) {
	user, ok := utils.GetUserFromContext(h.ctx)
	if !ok {
		return h.ctx, fmt.Errorf("user not found in context")
	}

	_, err := h.db.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		return h.ctx, fmt.Errorf("failed to delete account: %w", err)
	}

	h.ctx = context.Background()
	return h.ctx, nil
}

func (h *handler) ReportBuyerSpending() error {
	user, ok := utils.GetUserFromContext(h.ctx)
	if !ok {
		return fmt.Errorf("user not found in context")
	}

	fmt.Println("Buyer Spending Report:")

	query := `
    SELECT 
        o.id AS order_id, 
        u.id AS user_id, 
        u.full_name, 
        SUM(ci.qty * ci.price_at_purchase) AS total_spending, 
        SUM(ci.qty) AS total_qty
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
        AND u.id = ? 
    GROUP BY 
        o.id, u.id
    ORDER BY 
        total_spending DESC;
    `

	rows, err := h.db.Query(query, user.ID)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("\n+----------------+-----------------------+-----------------+------------+")
	fmt.Println("| Order ID      | User ID               | Full Name             | Total Spending  | Total Qty |")
	fmt.Println("+----------------+-----------------------+-----------------+------------+")

	var totalAmount float64
	for rows.Next() {
		var order_id int
		var user_id int
		var full_name string
		var total_spending float64
		var total_qty int

		err = rows.Scan(&order_id, &user_id, &full_name, &total_spending, &total_qty)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("| %-14d | %-21d | %-21s | %-15.2f | %-10d |\n", order_id, user_id, full_name, total_spending, total_qty)

		totalAmount += total_spending
	}

	fmt.Println("+----------------+-----------------------+-----------------+------------+")
	fmt.Printf("\nTotal Amount for All Orders: %-15.2f\n", totalAmount)

	return nil
}

func (h *handler) ReportUserWithHighestSpending() error {
	fmt.Println("Report: User with the Highest Spending Based on Price")
	user, ok := utils.GetUserFromContext(h.ctx)
	if !ok {
		return fmt.Errorf("user not found in context")
	}

	query := `
		SELECT 
			u.id AS user_id, 
			u.full_name, 
			SUM(ci.qty * ci.price_at_purchase) AS total_spending
		FROM 
			users u
		JOIN 
			carts c ON u.id = c.user_id
		JOIN 
			cart_items ci ON c.id = ci.user_id
		WHERE 
			c.status = 'Checked Out'
			AND u.toko_id = ?  -- Filter by toko_id
		GROUP BY 
			u.id
		ORDER BY 
			total_spending DESC
		LIMIT 10;
	`

	rows, err := h.db.Query(query, user.TokoID)

	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	fmt.Println("\n+----------------+-----------------------+-----------------+------------+")
	fmt.Println("| User ID       | Full Name             | Total Spending  |")
	fmt.Println("+----------------+-----------------------+-----------------+")

	for rows.Next() {
		var user_id int
		var full_name string
		var total_spending float64

		err = rows.Scan(&user_id, &full_name, &total_spending)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("| %-14d | %-21s | %-15.2f |\n", user_id, full_name, total_spending)
	}

	fmt.Println("+----------------+-----------------------+-----------------+")

	return nil
}

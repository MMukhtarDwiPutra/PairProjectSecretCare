package cli

import (
	"SecretCare/utils"
	"fmt"
)

func (c *cli) ReportBuyerSpending() {
	user, ok := utils.GetUserFromContext(c.ctx)
	if !ok {
		fmt.Print("user not found in context")
	}

	done := make(chan bool)
	go utils.LoadingSpinner(done)

	spendingReports, _ := c.handler.User.ReportBuyerSpending(user.ID)

	done <- true
	fmt.Print("\r                \r")

	fmt.Println("\n+----------------+-----------------------+-----------------+------------+")
	fmt.Println("| Order ID      | User ID               | Full Name             | Total Spending  | Total Qty |")
	fmt.Println("+----------------+-----------------------+-----------------+------------+")

	var totalAmount float64
	for _, spendingReport := range spendingReports {
		totalAmount += spendingReport.TotalSpending
		fmt.Printf("| %-14d | %-21d | %-21s | %-15.2f | %-10d |\n", spendingReport.OrderID, spendingReport.UserID, spendingReport.FullName, spendingReport.TotalSpending, spendingReport.TotalQuantity)
	}

	fmt.Println("+----------------+-----------------------+-----------------+------------+")
	fmt.Printf("\nTotal Amount for All Orders: %-15.2f\n", totalAmount)
}

func (c *cli) ReportUserWithHighestSpending() {
	user, ok := utils.GetUserFromContext(c.ctx)
	if !ok {
		fmt.Print("user not found in context")
	}

	done := make(chan bool)
	go utils.LoadingSpinner(done)

	userHighestSpending, _ := c.handler.User.ReportUserWithHighestSpending(user.ID)

	done <- true
	fmt.Print("\r                \r")

	fmt.Println("\n+----------------+-----------------------+-----------------+------------+")
	fmt.Println("| User ID       | Full Name             | Total Spending  |")
	fmt.Println("+----------------+-----------------------+-----------------+")

	var totalAmount float64
	for _, spendingReport := range userHighestSpending {
		totalAmount += spendingReport.TotalSpending
		fmt.Printf("| %-14d | %-21s | %-2.2f |\n", spendingReport.UserId, spendingReport.FullName, spendingReport.TotalSpending)
	}

	fmt.Println("+----------------+-----------------------+-----------------+------------+")
}

package entity

type Users struct {
	ID       int
	Username string
	Password string
	FullName string
	TokoID   int
	Role     string
}

type UserBuyerReport struct {
	OrderID       int
	UserID        int
	FullName      string
	TotalSpending float64
	TotalQuantity int
}

type UserReportHighestSpending struct {
	UserId        int
	FullName      string
	TotalSpending float64
}

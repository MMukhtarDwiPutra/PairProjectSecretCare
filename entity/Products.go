package entity

type Product struct{
	ID int
	Nama string
	Harga float64
	Stock int
	TokoID int
}

type ProductReport struct{
	Nama string
	TotalPenjualan int
	TotalPendapatan float64
}
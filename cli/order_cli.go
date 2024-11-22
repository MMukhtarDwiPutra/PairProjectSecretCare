package cli

import (
	"SecretCare/utils"
	"SecretCare/helpers"
	"fmt"
)	

func (c *cli) UpdateOrderStatus() {
	// add cart functionality
	user, _ := utils.GetUserFromContext(c.ctx)

	fmt.Println("Daftar Produk yang Tersedia:")
	orders, _ := c.handler.Order.GetAllOrderByTokoId(user.TokoID)

	if len(orders) == 0 {
		fmt.Println("Tidak ada orders tersedia.")
		return 
	}

	fmt.Println("")
	fmt.Println("================================")
	fmt.Println("Update Order Menjadi Checked Out")
	fmt.Println("================================")
	// Print the product list
	fmt.Println("ID Order\tNama Produk\tStatus Order")
	for _, order := range orders {
		fmt.Printf("%d\t%s\t%.2f\t%d\n", order.ID, order.NamaProduct, order.Status)
	}
	
	// masukan ID product
	orderID := helpers.InputAndHandlingNumber("Masukan id order yang ingin diubah statusnya: ")

	c.handler.Order.UpdateStatusOrder(orderID, "Checked Out")
}

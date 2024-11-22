package cli

import (
	"SecretCare/utils"
	"SecretCare/helpers"
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)	

// Buat reader agar bisa scan (input oleh user) nanti
var inputReader = bufio.NewReader(os.Stdin) 

func (c *cli) AddCart() {
	// add cart functionality
	user, _ := utils.GetUserFromContext(c.ctx)

	fmt.Println("Daftar Produk yang Tersedia:")
	products, _ := c.handler.Product.GetAllProducts()

	if len(products) == 0 {
		fmt.Println("Tidak ada produk tersedia.")
		return 
	}

	// Print the product list
	fmt.Println("ID\tNama Produk\tHarga\tStock")
	for _, product := range products {
		fmt.Printf("%d\t%s\t%.2f\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
	
	// masukan ID product
	fmt.Print("Masukan ID produk: ")
	productIDStr, _ := inputReader.ReadString('\n')
	productID, _ := strconv.Atoi(strings.TrimSpace(productIDStr))

	// masukan jumlah
	fmt.Print("Masukan jumlah: ")
	qtyStr, _ := inputReader.ReadString('\n')
	qty, _ := strconv.Atoi(strings.TrimSpace(qtyStr))

	fmt.Print("Masukan harga pembelian: ")
	priceStr, _ := inputReader.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

	c.handler.Cart.AddCart(user.ID, productID, qty, price)
}

func (c *cli) ShowCart() {
	// show Cart functionality
	user, _ := utils.GetUserFromContext(c.ctx)

	// Call the ShowCart function
	cartItems, _ := c.handler.Cart.ShowCart(user.ID)

	if len(cartItems) == 0 {
		fmt.Println("Keranjang Anda kosong.")
	} else {
		fmt.Println("Nama Product\tJumlah\tStatus Cart")
		for _, item := range cartItems {
			fmt.Printf("%s\t%d\t%s\n", item.ProductName, item.Quantity, item.Status)
		}
	}	
}

func (c *cli) Checkout(){
	// Checkout functionality
	user, _ := utils.GetUserFromContext(c.ctx)
	
	c.handler.Order.Checkout(user.ID)
}

func (c *cli) DeleteCart(){
	user, _ := utils.GetUserFromContext(c.ctx)
			
	// delete cart loop
	DeleteCartLoop: 
		for {
			fmt.Println("Pilih opsi:")
			fmt.Println("1. Delete All Cart Items")
			fmt.Println("2. Delete Satu Satu Barang")
			fmt.Println("3. Back")
			inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

			switch inputMenu {
				case 1:
					// Delete all items
					err := c.handler.Cart.DeleteAllCartItemsActive(user.ID)
					if err != nil {
						fmt.Printf("Gagal menghapus semua item di keranjang: %v\n", err)
					} else {
						fmt.Println("Berhasil menghapus semua item di keranjang.")
					}
				case 2:
					// Delete one item
					cartItems, err := c.handler.Cart.GetActiveCartItems(user.ID)
					if err != nil {
						fmt.Printf("Gagal mendapatkan item di keranjang: %v\n", err)
						continue
					}

					if len(cartItems) == 0 {
						fmt.Println("Tidak ada item di keranjang.")
						continue
					}

					// Print the cart items
					fmt.Println("ID\tNama Barang\tQty")
					for _, item := range cartItems {
						fmt.Printf("%d\t%s\t%d\n", item.ID, item.ProductName, item.Quantity)
					}

					// Get the cart item ID to delete
					cartItemID := helpers.InputAndHandlingNumber("Masukan ID item yang ingin dihapus: ")
					err = c.handler.Cart.DeleteCartItemByID(cartItemID)
					if err != nil {
						fmt.Printf("Gagal menghapus item dengan ID %d: %v\n", cartItemID, err)
					} else {
						fmt.Println("Berhasil menghapus item dari keranjang.")
					}
				case 3:
					// Back to previous menu
					break DeleteCartLoop
				default:
					fmt.Println("Masukan nomor menu yang valid.")
			}
		}	
}

func (c *cli) UpdateCart(){
	// Update Cart functionality
	user, _ := utils.GetUserFromContext(c.ctx)

	// Fetch active cart items for the user
	cartItems, _ := c.handler.Cart.GetActiveCartItems(user.ID)

	if len(cartItems) == 0 {
		fmt.Println("Tidak ada item di keranjang.")
		return
	}

	// Print the cart items
	fmt.Println("ID\tNama Barang\tQty")
	for _, item := range cartItems {
		fmt.Printf("%d\t%s\t%d\n", item.ID, item.ProductName, item.Quantity)
	}

	// Get the cart item ID to update
	cartItemID := helpers.InputAndHandlingNumber("Masukan ID item yang ingin diupdate: ")

	// Get the new quantity
	newQuantity := helpers.InputAndHandlingNumber("Masukan jumlah baru: ")

	// Update the cart item quantity
	c.handler.Cart.UpdateQuantityCart(cartItemID, newQuantity)
}
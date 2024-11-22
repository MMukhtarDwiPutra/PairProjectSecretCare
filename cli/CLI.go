package cli

import (
	// "ProjectDatabase/entity"
	"SecretCare/entity"
	"SecretCare/handler"
	"SecretCare/helpers"
	"SecretCare/utils"
	"bufio" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"context"
	"fmt"
	"os" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"strings"
	"strconv"
)

type CLI interface {
	Login(inputReader *bufio.Reader) (bool, string)
	Register(inputReader *bufio.Reader)
	MenuUtama()
}

type cli struct {
	handler *handler.Handler
	ctx     context.Context
}

func NewCli(handler *handler.Handler, ctx context.Context) *cli {
	return &cli{handler: handler, ctx: ctx}
}

func (c *cli) Login(inputReader *bufio.Reader) (bool, string) {
	var username, password string

	fmt.Print("Masukan username: ")
	username, _ = inputReader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukan password: ")
	password, _ = inputReader.ReadString('\n')
	password = strings.TrimSpace(password)

	successLogin, role, updatedCtx := c.handler.Auth.Login(username, password)


	c.ctx = updatedCtx
	return successLogin, role

}

func (c *cli) Register(inputReader *bufio.Reader) {
	var user entity.Users
	var inputRole int
	var namaToko string
	var toko entity.Toko

	user.TokoID = 1
	user.Role = ""

	for {
		fmt.Println("Anda ingin mendaftar sebagai apa?")
		fmt.Println("1. Penjual")
		fmt.Println("2. Pembeli")
		fmt.Println("3. Back")
		inputRole = helpers.InputAndHandlingNumber("Masukan input sesuai nomor yang ada diatas: ")

		if inputRole != 1 && inputRole != 2 && inputRole != 3 {
			fmt.Println("Masukan input sesuai nomor yang ada!")
			continue
		} else if inputRole == 1 {
			// Create toko
			fmt.Print("Masukan nama toko anda: ")
			namaToko, _ = inputReader.ReadString('\n')
			toko.Nama = strings.TrimSpace(namaToko)

			toko.ID = int(c.handler.Toko.CreateToko(context.Background(), toko))

			user.TokoID = toko.ID
			user.Role = "Penjual"
		} else if inputRole == 2 {
			user.Role = "Pembeli"
		} else if inputRole == 3 {
			fmt.Println("")
			return
		}

		break
	}

	for {
		// Input Full Name
		fmt.Print("Masukan nama panjang: ")
		user.FullName, _ = inputReader.ReadString('\n')
		user.FullName = strings.TrimSpace(user.FullName)

		// Input Username
		fmt.Print("Masukan username: ")
		user.Username, _ = inputReader.ReadString('\n')
		user.Username = strings.TrimSpace(user.Username)

		// Input Password
		fmt.Print("Masukan password: ")
		user.Password, _ = inputReader.ReadString('\n')
		user.Password = strings.TrimSpace(user.Password)

		// Confirm Password
		fmt.Print("Masukan confirm password: ")
		confirmPassword, _ := inputReader.ReadString('\n')
		confirmPassword = strings.TrimSpace(confirmPassword)

		// Check if passwords match
		if user.Password == confirmPassword {
			c.handler.Auth.RegisterUser(context.Background(), user)
			fmt.Println("Akun berhasil dibuat!")
			break
		} else {
			fmt.Println("Password dan confirm password yang dimasukan tidak sama!")
			continue
		}
	}
	fmt.Println("")
}

func (c *cli) UpdateMyAccount() {
	var username, password, fullName *string

	userInput := helpers.InputAndHandlingText("Masukan username baru (atau tekan Enter untuk melewati): ")
	if userInput != "" {
		username = &userInput
	}

	passwordInput := helpers.InputAndHandlingText("Masukan password baru (atau tekan Enter untuk melewati): ")
	if passwordInput != "" {
		password = &passwordInput
	}

	fullNameInput := helpers.InputAndHandlingText("Masukan nama lengkap baru (atau tekan Enter untuk melewati): ")
	if fullNameInput != "" {
		fullName = &fullNameInput
	}

	ctx, err := c.handler.User.UpdateMyAccount(username, password, fullName)
	c.ctx = ctx
	if err != nil {
		fmt.Printf("Gagal mengubah data akun: %v\n", err)
		return
	}

	updatedUser, ok := utils.GetUserFromContext(c.ctx)
	if !ok {
		fmt.Println("Tidak dapat mengambil data akun yang diperbarui.")
		return
	}

	fmt.Println("Data akun berhasil diubah.")
	fmt.Printf("Informasi akun terbaru:\nUsername: %s\nNama Lengkap: %s\n", updatedUser.Username, updatedUser.FullName)
}

func (c *cli) MenuProductReport(){
	user, _ := utils.GetUserFromContext(c.ctx)

	fmt.Println("")
	fmt.Println("==========================================================")
	fmt.Println("=====================Report Product=======================")
	fmt.Println("==========================================================")
	fmt.Printf("%-25s %-10v %s\n", "Nama Produk", "Penjualan", "Pendapatan") // Header with fixed column widths

	productReports := c.handler.Product.GetProductReport(user.TokoID)
	
	for _, productReport := range productReports{
		fmt.Printf("%-25s %-10v %.2f\n", productReport.Nama, productReport.TotalPenjualan, productReport.TotalPendapatan) // Header with fixed column widths
	}
}

func (c *cli) MenuUpdateStock(){
	user, _ := utils.GetUserFromContext(c.ctx)

	products := c.handler.Product.GetProductsByTokoID(user.TokoID)

	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Update Stock dari Produk")
	fmt.Println("=============================")
	fmt.Printf("%-5s %-25s %s\n", "ID", "Nama Produk", "Stock") // Header with fixed column widths

	for _, product := range products {
		// Print each product with aligned columns
		fmt.Printf("%-5d %-25s %d\n", product.ID, product.Nama, product.Stock)
	}

	fmt.Println("0. Untuk kembali")
	produkID := helpers.InputAndHandlingNumber("Masukan ID product yang ingin diupdate: ")
	if produkID == 0 {
		return
	}

	stock := helpers.InputAndHandlingNumber("Masukan jumlah stock terbaru: ")

	c.handler.Product.UpdateStockById(produkID, stock)
}

func (c *cli) MenuDeleteProduct(){
	user, _ := utils.GetUserFromContext(c.ctx)

	products := c.handler.Product.GetProductsByTokoID(user.TokoID)
	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Delete Produk dari Toko")
	fmt.Println("=============================")
	fmt.Println("ID\t Nama Produk")

	for _, product := range products{
		fmt.Printf("%v\t %v\n",product.ID, product.Nama)
	}
	fmt.Println("0. Untuk kembali")
	produkID := helpers.InputAndHandlingNumber("Masukan ID product yang ingin dihapus: ")
	if(produkID == 0){
		return
	}

	c.handler.Product.DeleteProductById(produkID)
	fmt.Println("Berhasil dihapus!")
}

func (c *cli) MenuCreateNewProduct(){
	var product entity.Product
	user, _ := utils.GetUserFromContext(c.ctx)

	inputReader := bufio.NewReader(os.Stdin) //Buat reader agar bisa scan (input oleh user) nanti

	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Tambah Product Baru di Toko")
	fmt.Println("=============================")
	// Input Username
	fmt.Print("Masukan nama produk: ")
	product.Nama, _ = inputReader.ReadString('\n')
	product.Nama = strings.TrimSpace(product.Nama)	

	// Input Username
	fmt.Print("Masukan harga: ")
	hargaString, _ := inputReader.ReadString('\n')
	hargaString = strings.TrimSpace(hargaString)
	product.Harga, _ = strconv.ParseFloat(hargaString, 64);

	// Input Username
	fmt.Print("Masukan stock: ")
	stockString, _ := inputReader.ReadString('\n')
	stockString = strings.TrimSpace(stockString)
	product.Stock, _ = strconv.Atoi(stockString);

	product.TokoID = user.TokoID
	c.handler.Product.CreateNewProduct(product);
}

func (c *cli) MenuPenjual() {
	var selesaiMenu bool = false

	users, _ := utils.GetUserFromContext(c.ctx)

	fmt.Print("Hallo mas " + users.FullName)

	for !selesaiMenu {
		fmt.Println("\nUser Penjual Menu")
		fmt.Println("1. Order Report")                   // Intermediate sql
		fmt.Println("2. Product Report")                 // Intermediate sql
		fmt.Println("3. Create New Produk untuk dijual") // easy golang
		fmt.Println("4. Update stock barang")            // intermediate golang
		fmt.Println("5. Delete barang")                  // easy sql
		fmt.Println("6. Akun")
		fmt.Println("7. Logout")
		inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu{
			case 1:
        		c.handler.User.ReportUserWithHighestSpending()
			case 2:
				c.MenuProductReport()
			case 3:
				c.MenuCreateNewProduct()
			case 4:
				c.MenuUpdateStock()
			case 5:
				c.MenuDeleteProduct()
			case 6:
				c.MenuAkun()
			case 7:
				selesaiMenu = true
		}

		fmt.Println()
	}
}

func (c *cli) MenuPembeli() {
	var selesaiMenu bool = false
	inputReader := bufio.NewReader(os.Stdin)

	for !selesaiMenu {
		fmt.Println("User Pembeli Menu")
		fmt.Println("1. Add Cart")    // Intermediate sql golang
		fmt.Println("2. Show Cart")   // Easy sql
		fmt.Println("3. Checkout")    // Hard beban kerja, sql golang
		fmt.Println("4. Delete Cart") // easy sql
		fmt.Println("5. Update Cart") // intermediate sql golang
		fmt.Println("6. Akun")
		fmt.Println("7. Buyer spending report")
		fmt.Println("8. Logout")
		inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu {
		case 1:
			// add cart functionality
			user, _ := utils.GetUserFromContext(c.ctx)

			fmt.Println("Daftar Produk yang Tersedia:")
   			products, _ := c.handler.Product.GetAllProducts()

			if len(products) == 0 {
				fmt.Println("Tidak ada produk tersedia.")
				continue
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
		case 2:
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
		case 3:
			// Checkout functionality
			user, _ := utils.GetUserFromContext(c.ctx)
			
			c.handler.Order.Checkout(user.ID)
		case 4:
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
		case 5:
			// Update Cart functionality
			user, _ := utils.GetUserFromContext(c.ctx)
		
			// Fetch active cart items for the user
			cartItems, _ := c.handler.Cart.GetActiveCartItems(user.ID)
		
			if len(cartItems) == 0 {
				fmt.Println("Tidak ada item di keranjang.")
				continue
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
		case 6:
			c.MenuAkun()
		case 7:
			c.handler.User.ReportBuyerSpending()
		case 8:
			selesaiMenu = true
		}

		fmt.Println()
	}
}

func (c *cli) MenuAkun() {
	var selesaiMenu bool = false

	for !selesaiMenu {
		fmt.Println("User Menu")
		fmt.Println("1. Delete Akun")    // easy sql
		fmt.Println("2. Ubah Data Akun") // intermediate golang
		fmt.Println("3. Back")
		inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu {
		case 1:
			ctx, _ := c.handler.User.DeleteMyAccount()
			c.ctx = ctx
			c.MenuUtama()
		case 2:
			c.UpdateMyAccount()
		case 3:
			selesaiMenu = true
		}

		fmt.Println()
	}
}

func (c *cli) MenuUtama() {
	var inputMenu int
	var selesaiMenu bool = false
	inputReader := bufio.NewReader(os.Stdin) //Buat reader agar bisa scan (input oleh user) nanti
	for !selesaiMenu {
		fmt.Println("User Menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		inputMenu = helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu {
		// Login
		case 1:
			successLogin, role := c.Login(inputReader)

			if successLogin {
				fmt.Println("Berhasil login!")

				switch role {
				case "Penjual":
					c.MenuPenjual()
				case "Pembeli":
					c.MenuPembeli()
				}
			} else {
				fmt.Println("Tidak berhasil login!")
			}

		// Register
		case 2:
			c.Register(inputReader)
		case 3:
			fmt.Println("Terima kasih telah menggunakan marketplace kami!")
			selesaiMenu = true
		}

		fmt.Println()
	}
}

package cli

import (
	"SecretCare/handler"
	"SecretCare/helpers"
	"SecretCare/utils"
	"context"
	"fmt"
	"bufio" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"os" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
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

		switch inputMenu {
		case 1:
			c.ReportUserWithHighestSpending()
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
			c.AddCart()	
		case 2:
			c.ShowCart()	
		case 3:
			c.Checkout()
		case 4:
			c.DeleteCart()	
		case 5:
			c.UpdateCart()
		case 6:
			c.MenuAkun()
		case 7:
			c.ReportBuyerSpending()
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
			c.DeleteMyAccount()
			c.ctx = context.Background()
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
				fmt.Println("\nLogged in!")
				fmt.Println("")

				switch role {
				case "Penjual":
					c.MenuPenjual()
				case "Pembeli":
					c.MenuPembeli()
				}
			} else {
				fmt.Println("\nUsername atau password salah!")
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

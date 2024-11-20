package cli

import(
	// "ProjectDatabase/entity"
	"fmt"
	// "bufio" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	// "os" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"SecretCare/helpers"
	"SecretCare/handler"
	// "SecretCare/entity"
	// "strings"
)

type CLI interface{
	MenuUtama()
}

type cli struct{
	handler handler.Handler
}

func NewCli(handler handler.Handler) *cli{
	return &cli{handler}
}

func (c *cli) MenuPenjual(){
	var selesaiMenu bool = false

	for !selesaiMenu{
		fmt.Println("User Menu")
		fmt.Println("1. Order Report") // Intermediate sql
		fmt.Println("2. Product Report") // Intermediate sql
		fmt.Println("3. Create New Produk untuk dijual") // easy golang
		fmt.Println("4. Update stock barang") // intermediate golang
		fmt.Println("5. Delete barang") // easy sql
		fmt.Println("6. Akun")
		fmt.Println("7. Logout")
		inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu{
			case 1:
			case 2:
			case 3:
			case 4:
			case 5:
			case 6:
				c.MenuAkun()
			case 7:
				selesaiMenu = true
		}

		fmt.Println()
	}
}

func (c *cli) MenuPembeli(){
	var selesaiMenu bool = false

	for !selesaiMenu{
		fmt.Println("User Menu")
		fmt.Println("1. Add Cart") // Intermediate sql golang
		fmt.Println("2. Show Cart") // Easy sql
		fmt.Println("3. Checkout") // Hard beban kerja, sql golang
		fmt.Println("4. Delete Cart") // easy sql
		fmt.Println("5. Update Cart") // intermediate sql golang
		fmt.Println("6. Akun")
		fmt.Println("7. Logout")
		inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu{
			case 1:
			case 2:
			case 3:
			case 4:
			case 5:
			case 6:
				c.MenuAkun()
			case 7:
				selesaiMenu = true
		}

		fmt.Println()
	}
}

func (c *cli) MenuAkun(){
	var selesaiMenu bool = false

	for !selesaiMenu{
		fmt.Println("User Menu")
		fmt.Println("1. Delete Akun") // easy sql
		fmt.Println("2. Ubah Data Akun") // intermediate golang
		fmt.Println("3. Back")
		inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu{
			case 1:
			case 2:
			case 3:
				selesaiMenu = true
		}

		fmt.Println()
	}
}

func (c *cli) MenuUtama(){
	var inputMenu int
	var selesaiMenu bool = false

	for !selesaiMenu{
		fmt.Println("User Menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		inputMenu = helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu{
			// Login
			case 1:
			// Register
			case 2:
			case 3:
				fmt.Println("Terima kasih telah menggunakan marketplace kami!")
				selesaiMenu = true
		}

		fmt.Println()
	}
}
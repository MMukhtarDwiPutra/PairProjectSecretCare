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
)

type CLI interface {
	Login(inputReader *bufio.Reader) (bool, string)
	Register(inputReader *bufio.Reader)
	MenuUtama()
}

type cli struct {
	handler handler.Handler
	ctx     context.Context
}

func NewCli(handler handler.Handler, ctx context.Context) *cli {
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

	successLogin, role, updatedCtx := c.handler.Handler.Login(username, password)

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

			toko.ID = int(c.handler.Handler.CreateToko(context.Background(), toko))

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
			c.handler.Handler.RegisterUser(context.Background(), user)
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

	ctx, err := c.handler.Handler.UpdateMyAccount(username, password, fullName)
	c.ctx = ctx
	if err != nil {
		fmt.Printf("Gagal mengubah data akun: %v\n", err)
	} else {
		fmt.Println("Data akun berhasil diubah")
	}
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
			c.handler.Handler.GetUserByUsername(users.Username)
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
		fmt.Println("7. Logout")
		inputMenu := helpers.InputAndHandlingNumber("Masukan nomor menu yang ingin dipilih: ")

		switch inputMenu {
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

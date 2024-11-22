package cli

import(
	"SecretCare/helpers"
	"SecretCare/entity"
	"context"
	"fmt"
	"strings"
	"bufio" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"SecretCare/utils"
)

func (c *cli) Login(inputReader *bufio.Reader) (bool, string) {
	var username, password string

	fmt.Print("Masukan username: ")
	username, _ = inputReader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukan password: ")
	password, _ = inputReader.ReadString('\n')
	password = strings.TrimSpace(password)

	done := make(chan bool)
	go utils.LoadingSpinner(done)

	successLogin, role, updatedCtx, _ := c.handler.Auth.Login(username, password)
	done <- true
	fmt.Print("\r                \r")

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

		for{
			// Input Username
			fmt.Print("Masukan username: ")
			user.Username, _ = inputReader.ReadString('\n')
			user.Username = strings.TrimSpace(user.Username)

			userTmp, _ := c.handler.User.GetUserByUsername(user.Username)
			if(userTmp != nil){
				fmt.Println("Username telah terdaftar! Silahkan input username lain!")
				continue
			}

			break
		}

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
			done := make(chan bool)
			go utils.LoadingSpinner(done)
			
			c.handler.Auth.RegisterUser(context.Background(), user)
			done <- true
			fmt.Print("\r                \r")
			fmt.Println("Akun berhasil dibuat!")
			break
		} else {
			fmt.Println("Password dan confirm password yang dimasukan tidak sama!")
			continue
		}
	}
	fmt.Println("")
}
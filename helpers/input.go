package helpers

import(
	"fmt" //Agar bisa memakai fungsi print di golang
	"bufio" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"os" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"strings" //Untuk mempermudah split string nanti (digunakan saat pengecekan multiple string)
	"strconv" //Untuk convert string ke integer nanti
)

// Membuat helpers fungsi untuk menghandle error jika inputan bukan angka
// Menerima input pesan dengan string untuk print ke layar cli nanti
// Mengembalikan nilai integer yang sudah diinput user
func InputAndHandlingNumber(pesan string) int{
	var intInput int = 0 //Untuk menampung input angka yang diinput oleh user
	inputReader := bufio.NewReader(os.Stdin) //Buat reader agar bisa scan (input oleh user) nanti

	for{
		fmt.Print(pesan) //Menampilkan string pesan yang sudah dimasukan dari parameter
		input, err := inputReader.ReadString('\n') //Input dimasukan ke variable input string terlebih dahulu untuk pengecekan string kosong
		input = strings.TrimSpace(input) //Hilangkan spasi pada variable input
		intInput, err = strconv.Atoi(input); //Cek jika bukan integer, maka ada error nantinya

		// Jika input bukan integer, maka program akan terus memaksa agar user menginput angka
		if (input == "" || err != nil) {
			fmt.Println("Invalid integer input!")
			continue //Skip kode setelah ini, dan balik ke loop awal
		}

		break //Looping forever berhenti
	}

	// Mengembalikan nilai integer yang sudah diinput user
	return intInput
}
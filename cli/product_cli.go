package cli

import(
	"fmt"
	"SecretCare/entity"
	"SecretCare/utils"
	"SecretCare/helpers"
	"os" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"bufio" //Diimpor untuk bisa scan multiple string dari layar cli di program golang
	"strings"
)

func (c *cli) MenuCreateNewProduct() {
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
	product.Harga = float64(helpers.InputAndHandlingNumber("Masukan harga: "))

	// Input Username
	product.Stock = helpers.InputAndHandlingNumber("Masukan stock: ")

	product.TokoID = user.TokoID
	_ = c.handler.Product.CreateNewProduct(product)
}

func (c *cli) MenuProductReport() {
	user, _ := utils.GetUserFromContext(c.ctx)

	fmt.Println("")
	fmt.Println("==========================================================")
	fmt.Println("=====================Report Product=======================")
	fmt.Println("==========================================================")
	fmt.Printf("%-25s %-10v %s\n", "Nama Produk", "Penjualan", "Pendapatan") // Header with fixed column widths

	done := make(chan bool)
	go utils.LoadingSpinner(done)

	productReports := c.handler.Product.GetProductReport(user.TokoID)
	done <- true
	fmt.Print("\r                \r")

	for _, productReport := range productReports {
		fmt.Printf("%-25s %-10v %.2f\n", productReport.Nama, productReport.TotalPenjualan, productReport.TotalPendapatan) // Header with fixed column widths
	}
}

func (c *cli) MenuUpdateStock() {
	user, _ := utils.GetUserFromContext(c.ctx)

	products := c.handler.Product.GetProductsByTokoID(user.TokoID)

	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Update Stock dari Produk")
	fmt.Println("=============================")
	fmt.Printf("%-5v %-25s %-5v %-5v\n", "ID", "Nama Produk", "Stock", "Harga") // Header with fixed column widths

	for _, product := range products {
		// Print each product with aligned columns
		fmt.Printf("%-5v %-25s %-5v %-5v\n", product.ID, product.Nama, product.Stock, product.Harga)
	}

	fmt.Println("\n0. Untuk kembali")
	produkID := helpers.InputAndHandlingNumber("Masukan ID product yang ingin diupdate: ")
	if produkID == 0 {
		return
	}

	fmt.Println("Update Stock %v")
	fmt.Println("")

	stock := helpers.InputAndHandlingNumber("Masukan jumlah stock terbaru: ")

	_ = c.handler.Product.UpdateStockById(produkID, stock)
}

func (c *cli) MenuDeleteProduct() {
	user, _ := utils.GetUserFromContext(c.ctx)

	products := c.handler.Product.GetProductsByTokoID(user.TokoID)
	fmt.Println("")
	fmt.Println("=============================")
	fmt.Println("Delete Produk dari Toko")
	fmt.Println("=============================")
	fmt.Println("ID\t Nama Produk")

	for _, product := range products {
		fmt.Printf("%v\t %v\n", product.ID, product.Nama)
	}
	fmt.Println("\n0. Untuk kembali")
	produkID := helpers.InputAndHandlingNumber("Masukan ID product yang ingin dihapus: ")
	if produkID == 0 {
		return
	}

	_ = c.handler.Product.DeleteProductById(produkID)
	fmt.Println("Berhasil dihapus!")
}
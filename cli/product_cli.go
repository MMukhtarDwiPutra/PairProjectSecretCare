package cli

import(
	"fmt"
	"SecretCare/entity"
	"SecretCare/utils"
	"SecretCare/helpers"
	"strconv"
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
	fmt.Print("Masukan harga: ")
	hargaString, _ := inputReader.ReadString('\n')
	hargaString = strings.TrimSpace(hargaString)
	product.Harga, _ = strconv.ParseFloat(hargaString, 64)

	// Input Username
	fmt.Print("Masukan stock: ")
	stockString, _ := inputReader.ReadString('\n')
	stockString = strings.TrimSpace(stockString)
	product.Stock, _ = strconv.Atoi(stockString)

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

	productReports := c.handler.Product.GetProductReport(user.TokoID)

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
	fmt.Println("0. Untuk kembali")
	produkID := helpers.InputAndHandlingNumber("Masukan ID product yang ingin dihapus: ")
	if produkID == 0 {
		return
	}

	_ = c.handler.Product.DeleteProductById(produkID)
	fmt.Println("Berhasil dihapus!")
}
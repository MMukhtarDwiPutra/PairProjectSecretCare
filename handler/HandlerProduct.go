package handler

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"SecretCare/entity"
)

type HandlerProduct interface{
	CreateNewProduct(product entity.Product)
	GetProductsByTokoID(tokoID int) []entity.Product
	DeleteProductById(id int)
	UpdateStockById(id int, stock int)
	GetProductReport(tokoID int) []entity.ProductReport
}

type handlerProduct struct{
	db *sql.DB
}

func NewHandlerProduct(db *sql.DB) *handlerProduct{
	return &handlerProduct{db}
}

func (h *handlerProduct) CreateNewProduct(product entity.Product){
	// Insert into the database
	_, err := h.db.Exec("INSERT INTO products (nama, harga, stock, toko_id) VALUES (?, ?, ?, ?)", product.Nama, product.Harga, product.Stock, product.TokoID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		fmt.Println()
		return
	}

	fmt.Println("Produk berhasil ditambahkan!")
}

func (h *handlerProduct) GetProductsByTokoID(tokoID int) []entity.Product{
	var products []entity.Product

	// Query the database
	rows, err := h.db.Query("SELECT id, nama, harga, stock, toko_id FROM products WHERE toko_id = ?", tokoID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return products
	}
	defer rows.Close() // Ensure rows are closed after use

	for rows.Next() {
		var product entity.Product

		// Scan the row into the product struct
		err := rows.Scan(&product.ID, &product.Nama, &product.Harga, &product.Stock, &product.TokoID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}

		products = append(products, product)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return products
}

func (h *handlerProduct) UpdateStockById(id int, stock int){
	_, err := h.db.Exec("UPDATE products SET stock = ? WHERE id = ?", stock, id)

	if err != nil {
		fmt.Println("Error executing query:", err)
		fmt.Println()
		return
	}
	fmt.Println("Stock product berhasil diupdate!")
}

func (h *handlerProduct) DeleteProductById(id int){
	_, err := h.db.Exec("DELETE FROM products WHERE id = ?", id)

	if err != nil {
		fmt.Println("Error executing query:", err)
		fmt.Println()
		return
	}
	fmt.Println("Product berhasil dihapus!")
}

func (h *handlerProduct) GetProductReport(tokoID int) []entity.ProductReport{
	var productReports []entity.ProductReport
	rows, err := h.db.Query(`SELECT 
							    products.nama, 
							    COALESCE(SUM(CASE WHEN orders.status = "Sudah dikirim" THEN cart_items.qty ELSE 0 END), 0) AS total_penjualan,
							    COALESCE(SUM(CASE WHEN orders.status = "Sudah dikirim" THEN cart_items.price_at_purchase ELSE 0 END), 0) AS total_pendapatan
							FROM 
							    products
							LEFT JOIN 
							    cart_items ON cart_items.product_id = products.id
							LEFT JOIN 
							    carts ON carts.id = cart_items.cart_id
							LEFT JOIN 
							    orders ON orders.cart_id = carts.id
							WHERE products.toko_id = ?
							GROUP BY 
							    products.nama;`, tokoID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return productReports
	}
	defer rows.Close() // Ensure rows are closed after use

	for rows.Next(){
		var productReport entity.ProductReport
		err := rows.Scan(&productReport.Nama, &productReport.TotalPenjualan, &productReport.TotalPendapatan)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}

		productReports = append(productReports, productReport)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return productReports
}
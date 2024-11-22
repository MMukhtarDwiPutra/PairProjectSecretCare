package handler

import (
	"SecretCare/entity"
	"context"
	"database/sql"
	"fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
)

type HandlerProduct interface {
	CreateNewProduct(product entity.Product) error
	GetProductsByTokoID(tokoID int) []entity.Product
	DeleteProductById(id int) error
	UpdateStockById(id int, stock int) error
	GetProductReport(tokoID int) []entity.ProductReport
	GetAllProducts() ([]entity.Product, error)
}

type handlerProduct struct {
	ctx context.Context
	db  *sql.DB
}

// NewHandlerProduct membuat instance baru dari HandlerProduct
func NewHandlerProduct(ctx context.Context, db *sql.DB) *handlerProduct {
	return &handlerProduct{ctx, db}
}

func (h *handlerProduct) CreateNewProduct(product entity.Product) error {
	query := `
		INSERT INTO products (nama, harga, stock, toko_id) 
		VALUES ($1, $2, $3, $4)
	`
	_, err := h.db.ExecContext(h.ctx, query, product.Nama, product.Harga, product.Stock, product.TokoID)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err)
	}

	fmt.Println("Produk berhasil ditambahkan!")
	return nil
}

func (h *handlerProduct) GetProductsByTokoID(tokoID int) []entity.Product {
	var products []entity.Product
	query := `
		SELECT id, nama, harga, stock, toko_id 
		FROM products 
		WHERE toko_id = $1
	`

	rows, err := h.db.QueryContext(h.ctx, query, tokoID)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return products
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Nama, &product.Harga, &product.Stock, &product.TokoID)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		products = append(products, product)
	}

	return products
}

func (h *handlerProduct) UpdateStockById(id int, stock int) error {
	query := `
		UPDATE products 
		SET stock = $1 
		WHERE id = $2
	`
	_, err := h.db.ExecContext(h.ctx, query, stock, id)
	if err != nil {
		return fmt.Errorf("failed to update stock: %v", err)
	}

	fmt.Println("Stock product berhasil diupdate!")
	return nil
}

func (h *handlerProduct) DeleteProductById(id int) error {
	query := `
		DELETE FROM products 
		WHERE id = $1
	`
	_, err := h.db.ExecContext(h.ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	fmt.Println("Product berhasil dihapus!")
	return nil
}

func (h *handlerProduct) GetProductReport(tokoID int) []entity.ProductReport {
	var productReports []entity.ProductReport
	query := `
		SELECT 
			products.nama, 
			COALESCE(SUM(CASE WHEN orders.status = 'Shipped' THEN cart_items.qty ELSE 0 END), 0) AS total_penjualan,
			COALESCE(SUM(CASE WHEN orders.status = 'Shipped' THEN cart_items.price_at_purchase ELSE 0 END), 0) AS total_pendapatan
		FROM 
			products
		LEFT JOIN 
			cart_items ON cart_items.product_id = products.id
		LEFT JOIN 
			carts ON carts.id = cart_items.cart_id
		LEFT JOIN 
			orders ON orders.cart_id = carts.id
		WHERE 
			products.toko_id = $1
		GROUP BY 
			products.nama
	`

	rows, err := h.db.QueryContext(h.ctx, query, tokoID)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return productReports
	}
	defer rows.Close()

	for rows.Next() {
		var productReport entity.ProductReport
		err := rows.Scan(&productReport.Nama, &productReport.TotalPenjualan, &productReport.TotalPendapatan)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		productReports = append(productReports, productReport)
	}

	return productReports
}

func (h *handlerProduct) GetAllProducts() ([]entity.Product, error) {
	query := `
		SELECT id, nama, harga, stock 
		FROM products
	`
	rows, err := h.db.QueryContext(h.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Nama, &product.Harga, &product.Stock)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

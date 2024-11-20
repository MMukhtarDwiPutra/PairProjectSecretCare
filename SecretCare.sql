-- Membuat database baru bernama SecretCare
CREATE DATABASE SecretCare;

CREATE TABLE toko(
	id INT PRIMARY KEY AUTO_INCREMENT,
	nama VARCHAR(256) NOT NULL
);

CREATE TABLE products(
	id INT PRIMARY KEY AUTO_INCREMENT,
	nama VARCHAR(256) NOT NULL,
	harga INT NOT NULL,
	stock INT DEFAULT 0,
	toko_id INT,
	FOREIGN KEY (toko_id) REFERENCES toko(id)
);

CREATE TABLE users(
	id INT PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(64) NOT NULL,
	password VARCHAR(64) NOT NULL,
	full_name VARCHAR(128) NOT NULL,
	toko_id INT DEFAULT 1,
	role ENUM('Penjual', 'Pembeli')
);

CREATE TABLE carts(
	id INT PRIMARY KEY AUTO_INCREMENT,
	status ENUM("Active", "Checked Out"),
	user_id INT,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE cart_items(
	id INT PRIMARY KEY AUTO_INCREMENT,
	user_id INT,
	product_id INT,
	qty INT NOT NULL,
	price_at_purchase FLOAT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (product_id) REFERENCES products(id)
);


CREATE TABLE orders(
	id INT PRIMARY KEY AUTO_INCREMENT,
	status VARCHAR(64) NOT NULL,
	order_date DATE NOT NULL,
	cart_id INT,
	FOREIGN KEY (cart_id) REFERENCES carts(id)
);

INSERT INTO toko(id, nama) VALUES
(1, "Tidak ada toko"),
(2, "Toko Serba Ada"),
(3, "Toko Bang Fathur");

INSERT INTO users(id, username, password, full_name, toko_id, role) VALUES
(1, "mmukhtar", "$2a$14$VQC3HrTWFV1THTB67iaQ1OfbsDD1lzBxBpYmXKmpXxSvhz6/dVmkC", "Muhammad Mukhtar Dwi Putra", 2, 1),
(2, "fathur", "$2a$14$TqS2QWIy3xxmbvJ9DFbLZOeeIpzTK5aRR45jrMkSfHEkDJ1t.VKCu", "Muhammad Mukhtar Dwi Putra", 3, 1),
(3, "obiea", "$2a$14$1dzedAhWAUYtFyt/ylDDC.98xmvt6OELiZrQ4C3ovxLeZw58qPHa6", "Muhammad Mukhtar Dwi Putra", 1, 2);

INSERT INTO products(id, nama, harga, stock, toko_id) VALUES
(1, "Pelembab Skintific Vit C", 50000, 23, 2),
(2, "Serum Skintific Vit C", 75000, 10, 2),
(3, "Toner Skintific Vit C", 60000, 12, 2),
(4, "Pelembab Wardah AHA/BHA", 50000, 34, 3),
(5, "Serum Wardah AHA/BHA", 50000, 39, 3);
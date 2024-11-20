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

INSERT INTO users(id, username, password, full_name, role_id, toko_id) VALUES
(1, "admin", "admin", "Administrator", 3, 1),
(2, "mmukhtar", "pwmmukhtar", "Muhammad Mukhtar", 1, 2),
(3, "fathur", "pwfathur", "Fathur Rohman", 1, 3),
(4, "obie", "pwobie", "Obie Ananda", 2, 1);

INSERT INTO toko(id, nama) VALUES
(1, "Tidak ada toko"),
(2, "Toko Mukhtar"),
(3, "Toko Fathur");

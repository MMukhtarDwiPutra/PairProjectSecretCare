-- Membuat database baru bernama SecretCare
CREATE DATABASE SecretCare;

USE SecretCare;

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
	cart_id INT,
	product_id INT,
	qty INT NOT NULL,
	price_at_purchase FLOAT NOT NULL,
	FOREIGN KEY (cart_id) REFERENCES carts(id),
	FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);


CREATE TABLE orders(
    id INT PRIMARY KEY AUTO_INCREMENT,
    status ENUM('Shipped', 'Waiting For Payment', 'Checked Out') NOT NULL,
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
(2, "fathur", "$2a$14$TqS2QWIy3xxmbvJ9DFbLZOeeIpzTK5aRR45jrMkSfHEkDJ1t.VKCu", "Muhammad Mukhtar Dwi Putra", 3, 2),
(3, "obiea", "$2a$14$1dzedAhWAUYtFyt/ylDDC.98xmvt6OELiZrQ4C3ovxLeZw58qPHa6", "Muhammad Mukhtar Dwi Putra", 1, 2);

INSERT INTO carts (status, user_id)
VALUES 
    ('Active', 2),
    ('Checked Out', 2),
    ('Active', 2);

INSERT INTO products(id, nama, harga, stock, toko_id) VALUES
(1, "Pelembab Skintific Vit C", 50000, 23, 2),
(2, "Serum Skintific Vit C", 75000, 10, 2),
(3, "Toner Skintific Vit C", 60000, 12, 2),
(4, "Pelembab Wardah AHA/BHA", 50000, 34, 3),
(5, "Serum Wardah AHA/BHA", 50000, 39, 3);

INSERT INTO `carts` (`id`, `status`, `user_id`) VALUES
(1, 'Checked Out', 3),
(2, 'Active', 3);

INSERT INTO `cart_items` (`id`, `cart_id`, `product_id`, `qty`, `price_at_purchase`) VALUES
(1, 1, 1, 10, 100000),
(2, 1, 2, 29, 290000),
(3, 2, 5, 21, 210000);

INSERT INTO `orders` (`id`, `status`, `order_date`, `cart_id`) VALUES
(1, 'Sudah Dikirim', '2024-11-05', 1),
(2, 'Belum dikirim', '2024-11-05', 2);

INSERT INTO cart_items (cart_id, product_id, qty, price_at_purchase)
VALUES 
    (7, 1, 2, 10000),
    (8, 2, 1, 15000),
    (9, 3, 1, 20000);

	INSERT INTO orders (status, order_date, cart_id)
VALUES 
    ('Waiting For Payment', '2024-11-01', 7),
    ('Checked Out', '2024-11-15', 8);
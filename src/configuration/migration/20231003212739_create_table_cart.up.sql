CREATE TABLE IF NOT EXISTS cart (
    cart_id VARCHAR(36) PRIMARY KEY,
    user_id BIGINT,
    product_id VARCHAR(36),
    product_name VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(uuid)
);
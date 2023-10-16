CREATE TABLE IF NOT EXISTS cart (
    cart_id VARCHAR(36) PRIMARY KEY,
    user_id BIGINT,
    quantity INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(uuid)
);
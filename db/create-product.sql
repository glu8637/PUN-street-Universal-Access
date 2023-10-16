CREATE TABLE IF NOT EXISTS product (
    product_id SERIAL PRIMARY KEY,
    storeId SERIAL references stores(store_id),
    name VARCHAR(255) NOT NULL,
    describe VARCHAR(255),
    category_id INT,
    picture BYTEA,
    price INT NOT NULL,
    stock INT NOT NULL
);

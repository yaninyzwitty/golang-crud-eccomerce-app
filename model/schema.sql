CREATE KEYSPACE IF NOT EXISTS merchandise_system WITH replication = { 'class': 'SimpleStrategy',
'replication_factor': 1 };

USE merchandise_system;

CREATE TABLE IF NOT EXISTS products (
    product_id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    price DOUBLE,
    quantity INT
);

CREATE TABLE IF NOT EXISTS categories (
    category_id UUID PRIMARY KEY,
    name TEXT,
    description TEXT
);

CREATE TABLE IF NOT EXISTS product_categories (
    product_id UUID,
    category_id UUID,
    PRIMARY KEY (product_id, category_id)
);

CREATE TABLE IF NOT EXISTS orders (
    order_id UUID PRIMARY KEY,
    customer_id UUID,
    order_date TIMESTAMP,
    total_amount DOUBLE
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id UUID,
    product_id UUID,
    quantity INT,
    price DOUBLE,
    PRIMARY KEY (order_id, product_id)
);

CREATE TABLE IF NOT EXISTS customers (
    customer_id UUID PRIMARY KEY,
    name TEXT,
    email TEXT,
    address TEXT
);

CREATE TABLE IF NOT EXISTS reviews (
    review_id UUID PRIMARY KEY,
    product_id UUID,
    customer_id UUID,
    rating INT,
    comment TEXT
);
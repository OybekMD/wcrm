CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(35),
    last_name VARCHAR(35),
    username VARCHAR(20),
    phone_number VARCHAR(15),
    bio VARCHAR(300),
    birth_day DATE,
    email VARCHAR(40),
    avatar TEXT, 
    password VARCHAR(75),
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE category_icons(
    id SERIAL PRIMARY KEY,
    name VARCHAR(45),
    picture TEXT
);

CREATE TABLE categorys(
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    icon_id INT REFERENCES category_icons(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(65),
    description VARCHAR(255),
    price INT,
    picture TEXT,
    category_id INT REFERENCES categorys(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE orderproducts (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    product_id INT REFERENCES products(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content VARCHAR(255),
    user_id UUID,
    product_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    email VARCHAR(40),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);




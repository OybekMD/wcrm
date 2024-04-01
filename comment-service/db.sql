CREATE DATABASE crmpostdb;
\c crmpostdb

CREATE TABLE category_icons(
    id SERIAL PRIMARY KEY,
    name VARCHAR(45),
    picture TEXT
);

CREATE TABLE categorys(
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    user_id VARCHAR(40),
    icon_id INT REFERENCES category_icons(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE product_images(
    id SERIAL PRIMARY KEY,
    picture TEXT,
    product_id INT REFERENCES products(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(65),
    description VARCHAR(255),
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE orderproducts (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(40),
    picture_id INT REFERENCES products(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE category_icons;
DROP TABLE product_images;
DROP TABLE categorys;
DROP TABLE products;
DROP TABLE orderproducts;

\c postgres
DROP DATABASE crmpostdb;





-- Media Update domain name
UPDATE languages
SET picture = REPLACE(picture, 'localhost', 'domain.com')
WHERE picture LIKE '%localhost%';

UPDATE languages
SET picture = REPLACE(picture, 'domain.com', 'localhost')
WHERE picture LIKE '%domain%';

UPDATE lessons
SET deleted_at = "";


-- Update
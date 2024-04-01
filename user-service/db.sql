CREATE DATABASE crmuserdb;
\c crmuserdb

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

DROP TABLE users;
DROP DATABASE crmuserdb;
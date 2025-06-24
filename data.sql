-- Active: 1750554601825@@127.0.0.1@5434@postgres
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

SELECT * FROM users
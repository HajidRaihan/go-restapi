CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255)
);
-- migrate -database "postgres://postgres:postgres@localhost:5432/go-gin-gonic?sslmode=disable" -path database/migrations up
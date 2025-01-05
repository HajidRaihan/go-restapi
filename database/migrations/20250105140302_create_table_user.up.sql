CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    born_date TIMESTAMP
);
-- migrate -database "postgres://postgres:postgres@localhost:5432/go-gin-gonic?sslmode=disable" -path database/migrations up

-- migrate create -ext sql -dir db/migrations -seq create_users_table
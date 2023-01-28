-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    age INTEGER NOT NULL,
    some_id INTEGER NOT NULL
);
-- +goose Down
DROP TABLE user;
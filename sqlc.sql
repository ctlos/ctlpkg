-- name: InsertUser :one

INSERT INTO users (name, age, description, some_id)
VALUES ($1, $2, $3, $4) RETURNING id;

-- name: SelectUsers :many

SELECT *
FROM users;

-- name: DeleteUser :exec

DELETE
FROM users
WHERE id = $1;

-- name: UpdateUser :exec

UPDATE users
SET name = $2,
    age = $3,
    description = $4,
    some_id = $5
WHERE id = $1;
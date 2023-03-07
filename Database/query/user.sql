-- name: CreateAddress :one
INSERT INTO address (
unit_number,
address_line1,
address_line2,
postal_code
) VALUES(
    $1, $2, $3, $4
) RETURNING *;

-- name: CreateUser :one
INSERT INTO user_details (
first_name,
last_name,
user_name,
user_password,
email,
mobile
) VALUES(
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM user_details WHERE user_name = $1;
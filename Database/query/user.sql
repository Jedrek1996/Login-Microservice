-- name: CreateAddress :one
INSERT INTO address (
unit_number,
address_line1,
address_line2,
postal_code
) VALUES(
    $1, $2, $3, $4
) RETURNING *;

-- name: CreateU2ser :one
INSERT INTO user_details (
first_name,
last_name,
user_name,
email,
mobile
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;
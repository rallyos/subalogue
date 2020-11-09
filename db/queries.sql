-- name: CreateUser :one
INSERT INTO users (username) VALUES($1) RETURNING *;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: CreateSubscription :one
INSERT INTO subscriptions (
    user_id, name, url, price
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: ListSubscriptions :many
SELECT * FROM subscriptions where user_id = $1;

-- name: GetSubscription :one
SELECT * FROM subscriptions where id = $1 AND user_id = $2;

-- name: UpdateSubscription :one
UPDATE subscriptions
SET name = $2, url=$3, price=$4
WHERE user_id = $1 AND id = $5 RETURNING *;

-- name: DeleteSubscription :one
DELETE FROM subscriptions WHERE user_id = $1 AND id = $2 RETURNING *;

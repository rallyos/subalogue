-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: CreateUserSubscription :one
INSERT INTO subscriptions (
    user_id, name, price
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: ListUserSubscriptions :many
SELECT * FROM subscriptions where user_id = $1;

-- name: UpdateUserSubscription :exec
UPDATE subscriptions
SET name = $2, price = $3
WHERE user_id = $1 AND id = $4;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions WHERE user_id = $1 AND id = $2;

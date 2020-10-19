-- name: CreateSubscriptions :one
INSERT INTO subscriptions (
    name, price
) VALUES (
    $1, $2
)
RETURNING *;

-- name: ListSubscriptions :many
SELECT * FROM subscriptions;

-- name: UpdateSubscriptions :exec
UPDATE subscriptions
SET name = $1, price = $2
WHERE id = $3;

-- name: DeleteSubscriptions :exec
DELETE FROM subscriptions WHERE id = $1;

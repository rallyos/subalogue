// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package db

import (
	"context"
)

const createUserSubscription = `-- name: CreateUserSubscription :one
INSERT INTO subscriptions (
    user_id, name, price
) VALUES (
    $1, $2, $3
)
RETURNING id, name, user_id, price
`

type CreateUserSubscriptionParams struct {
	UserID int32
	Name   string
	Price  int32
}

func (q *Queries) CreateUserSubscription(ctx context.Context, arg CreateUserSubscriptionParams) (Subscription, error) {
	row := q.db.QueryRowContext(ctx, createUserSubscription, arg.UserID, arg.Name, arg.Price)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Price,
	)
	return i, err
}

const deleteSubscription = `-- name: DeleteSubscription :exec
DELETE FROM subscriptions WHERE user_id = $1 AND id = $2
`

type DeleteSubscriptionParams struct {
	UserID int32
	ID     int32
}

func (q *Queries) DeleteSubscription(ctx context.Context, arg DeleteSubscriptionParams) error {
	_, err := q.db.ExecContext(ctx, deleteSubscription, arg.UserID, arg.ID)
	return err
}

const findUserByUsername = `-- name: FindUserByUsername :one
SELECT id, username FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) FindUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByUsername, username)
	var i User
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const listUserSubscriptions = `-- name: ListUserSubscriptions :many
SELECT id, name, user_id, price FROM subscriptions where user_id = $1
`

func (q *Queries) ListUserSubscriptions(ctx context.Context, userID int32) ([]Subscription, error) {
	rows, err := q.db.QueryContext(ctx, listUserSubscriptions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subscription
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserSubscription = `-- name: UpdateUserSubscription :exec
UPDATE subscriptions
SET name = $2, price = $3
WHERE user_id = $1 AND id = $4
`

type UpdateUserSubscriptionParams struct {
	UserID int32
	Name   string
	Price  int32
	ID     int32
}

func (q *Queries) UpdateUserSubscription(ctx context.Context, arg UpdateUserSubscriptionParams) error {
	_, err := q.db.ExecContext(ctx, updateUserSubscription,
		arg.UserID,
		arg.Name,
		arg.Price,
		arg.ID,
	)
	return err
}

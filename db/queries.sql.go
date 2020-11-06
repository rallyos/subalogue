// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package db

import (
	"context"
)

const createSubscription = `-- name: CreateSubscription :one
INSERT INTO subscriptions (
    user_id, name, url, price
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, name, user_id, price, url
`

type CreateSubscriptionParams struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Price  int32  `json:"price"`
}

func (q *Queries) CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) (Subscription, error) {
	row := q.db.QueryRowContext(ctx, createSubscription,
		arg.UserID,
		arg.Name,
		arg.Url,
		arg.Price,
	)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Price,
		&i.Url,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username) VALUES($1) RETURNING id, username
`

func (q *Queries) CreateUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, username)
	var i User
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const deleteSubscription = `-- name: DeleteSubscription :exec
DELETE FROM subscriptions WHERE user_id = $1 AND id = $2
`

type DeleteSubscriptionParams struct {
	UserID int32 `json:"user_id"`
	ID     int32 `json:"id"`
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

const getSubscription = `-- name: GetSubscription :one
SELECT id, name, user_id, price, url FROM subscriptions where id = $1 AND user_id = $2
`

type GetSubscriptionParams struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) GetSubscription(ctx context.Context, arg GetSubscriptionParams) (Subscription, error) {
	row := q.db.QueryRowContext(ctx, getSubscription, arg.ID, arg.UserID)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Price,
		&i.Url,
	)
	return i, err
}

const listSubscriptions = `-- name: ListSubscriptions :many
SELECT id, name, user_id, price, url FROM subscriptions where user_id = $1
`

func (q *Queries) ListSubscriptions(ctx context.Context, userID int32) ([]Subscription, error) {
	rows, err := q.db.QueryContext(ctx, listSubscriptions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Subscription{}
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Price,
			&i.Url,
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

const updateSubscription = `-- name: UpdateSubscription :one
UPDATE subscriptions
SET name = $2, url=$3, price=$4
WHERE user_id = $1 AND id = $5 RETURNING id, name, user_id, price, url
`

type UpdateSubscriptionParams struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Price  int32  `json:"price"`
	ID     int32  `json:"id"`
}

func (q *Queries) UpdateSubscription(ctx context.Context, arg UpdateSubscriptionParams) (Subscription, error) {
	row := q.db.QueryRowContext(ctx, updateSubscription,
		arg.UserID,
		arg.Name,
		arg.Url,
		arg.Price,
		arg.ID,
	)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Price,
		&i.Url,
	)
	return i, err
}

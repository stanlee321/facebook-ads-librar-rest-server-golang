// Code generated by sqlc. DO NOT EDIT.
// source: facebook_ad_demo.sql

package db

import (
	"context"
	"database/sql"
)

const createFacebookDemo = `-- name: CreateFacebookDemo :one
INSERT INTO "FacebookDemos" (
  ad_id,
  page_id,
  age,
  gender,
  percentage,
  ad_delivery_start_time
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, ad_id, page_id, age, gender, percentage, ad_delivery_start_time, created_at
`

type CreateFacebookDemoParams struct {
	AdID                sql.NullInt64  `json:"ad_id"`
	PageID              sql.NullInt64  `json:"page_id"`
	Age                 sql.NullInt32  `json:"age"`
	Gender              sql.NullString `json:"gender"`
	Percentage          interface{}    `json:"percentage"`
	AdDeliveryStartTime sql.NullString `json:"ad_delivery_start_time"`
}

func (q *Queries) CreateFacebookDemo(ctx context.Context, arg CreateFacebookDemoParams) (FacebookDemo, error) {
	row := q.db.QueryRowContext(ctx, createFacebookDemo,
		arg.AdID,
		arg.PageID,
		arg.Age,
		arg.Gender,
		arg.Percentage,
		arg.AdDeliveryStartTime,
	)
	var i FacebookDemo
	err := row.Scan(
		&i.ID,
		&i.AdID,
		&i.PageID,
		&i.Age,
		&i.Gender,
		&i.Percentage,
		&i.AdDeliveryStartTime,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFaceookDemo = `-- name: DeleteFaceookDemo :exec
DELETE FROM "FacebookDemos"
WHERE id = $1
`

func (q *Queries) DeleteFaceookDemo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFaceookDemo, id)
	return err
}

const getFacebookDemo = `-- name: GetFacebookDemo :one
SELECT id, ad_id, page_id, age, gender, percentage, ad_delivery_start_time, created_at FROM "FacebookDemos"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFacebookDemo(ctx context.Context, id int64) (FacebookDemo, error) {
	row := q.db.QueryRowContext(ctx, getFacebookDemo, id)
	var i FacebookDemo
	err := row.Scan(
		&i.ID,
		&i.AdID,
		&i.PageID,
		&i.Age,
		&i.Gender,
		&i.Percentage,
		&i.AdDeliveryStartTime,
		&i.CreatedAt,
	)
	return i, err
}

const listFacebookDemos = `-- name: ListFacebookDemos :many
SELECT id, ad_id, page_id, age, gender, percentage, ad_delivery_start_time, created_at FROM "FacebookDemos"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListFacebookDemosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFacebookDemos(ctx context.Context, arg ListFacebookDemosParams) ([]FacebookDemo, error) {
	rows, err := q.db.QueryContext(ctx, listFacebookDemos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FacebookDemo{}
	for rows.Next() {
		var i FacebookDemo
		if err := rows.Scan(
			&i.ID,
			&i.AdID,
			&i.PageID,
			&i.Age,
			&i.Gender,
			&i.Percentage,
			&i.AdDeliveryStartTime,
			&i.CreatedAt,
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

const listFacebookDemosByAdID = `-- name: ListFacebookDemosByAdID :many
SELECT id, ad_id, page_id, age, gender, percentage, ad_delivery_start_time, created_at FROM "FacebookDemos"
WHERE ad_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListFacebookDemosByAdIDParams struct {
	AdID   sql.NullInt64 `json:"ad_id"`
	Limit  int32         `json:"limit"`
	Offset int32         `json:"offset"`
}

func (q *Queries) ListFacebookDemosByAdID(ctx context.Context, arg ListFacebookDemosByAdIDParams) ([]FacebookDemo, error) {
	rows, err := q.db.QueryContext(ctx, listFacebookDemosByAdID, arg.AdID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FacebookDemo{}
	for rows.Next() {
		var i FacebookDemo
		if err := rows.Scan(
			&i.ID,
			&i.AdID,
			&i.PageID,
			&i.Age,
			&i.Gender,
			&i.Percentage,
			&i.AdDeliveryStartTime,
			&i.CreatedAt,
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

const listFacebookDemosByPageID = `-- name: ListFacebookDemosByPageID :many
SELECT id, ad_id, page_id, age, gender, percentage, ad_delivery_start_time, created_at FROM "FacebookDemos"
WHERE page_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListFacebookDemosByPageIDParams struct {
	PageID sql.NullInt64 `json:"page_id"`
	Limit  int32         `json:"limit"`
	Offset int32         `json:"offset"`
}

func (q *Queries) ListFacebookDemosByPageID(ctx context.Context, arg ListFacebookDemosByPageIDParams) ([]FacebookDemo, error) {
	rows, err := q.db.QueryContext(ctx, listFacebookDemosByPageID, arg.PageID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FacebookDemo{}
	for rows.Next() {
		var i FacebookDemo
		if err := rows.Scan(
			&i.ID,
			&i.AdID,
			&i.PageID,
			&i.Age,
			&i.Gender,
			&i.Percentage,
			&i.AdDeliveryStartTime,
			&i.CreatedAt,
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
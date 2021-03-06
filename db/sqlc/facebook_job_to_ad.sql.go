// Code generated by sqlc. DO NOT EDIT.
// source: facebook_job_to_ad.sql

package db

import (
	"context"
	"database/sql"
)

const createJobToFacebookAd = `-- name: CreateJobToFacebookAd :one
INSERT INTO "JobToFacebookAd" (
  job_id,
  ad_id
) VALUES (
  $1, $2
)
RETURNING id, job_id, ad_id, created_at
`

type CreateJobToFacebookAdParams struct {
	JobID sql.NullInt64 `json:"job_id"`
	AdID  sql.NullInt64 `json:"ad_id"`
}

func (q *Queries) CreateJobToFacebookAd(ctx context.Context, arg CreateJobToFacebookAdParams) (JobToFacebookAd, error) {
	row := q.db.QueryRowContext(ctx, createJobToFacebookAd, arg.JobID, arg.AdID)
	var i JobToFacebookAd
	err := row.Scan(
		&i.ID,
		&i.JobID,
		&i.AdID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteJobToFacebookAd = `-- name: DeleteJobToFacebookAd :exec
DELETE FROM "JobToFacebookAd"
WHERE id = $1
`

func (q *Queries) DeleteJobToFacebookAd(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteJobToFacebookAd, id)
	return err
}

const getJobToFacebookAd = `-- name: GetJobToFacebookAd :one
SELECT id, job_id, ad_id, created_at FROM "JobToFacebookAd"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetJobToFacebookAd(ctx context.Context, id int64) (JobToFacebookAd, error) {
	row := q.db.QueryRowContext(ctx, getJobToFacebookAd, id)
	var i JobToFacebookAd
	err := row.Scan(
		&i.ID,
		&i.JobID,
		&i.AdID,
		&i.CreatedAt,
	)
	return i, err
}

const listJobToFacebookAd = `-- name: ListJobToFacebookAd :many
SELECT id, job_id, ad_id, created_at FROM "JobToFacebookAd"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListJobToFacebookAdParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListJobToFacebookAd(ctx context.Context, arg ListJobToFacebookAdParams) ([]JobToFacebookAd, error) {
	rows, err := q.db.QueryContext(ctx, listJobToFacebookAd, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []JobToFacebookAd{}
	for rows.Next() {
		var i JobToFacebookAd
		if err := rows.Scan(
			&i.ID,
			&i.JobID,
			&i.AdID,
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

const listJobToFacebookAdByJobID = `-- name: ListJobToFacebookAdByJobID :many
SELECT id, job_id, ad_id, created_at FROM "JobToFacebookAd"
WHERE job_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListJobToFacebookAdByJobIDParams struct {
	JobID  sql.NullInt64 `json:"job_id"`
	Limit  int32         `json:"limit"`
	Offset int32         `json:"offset"`
}

func (q *Queries) ListJobToFacebookAdByJobID(ctx context.Context, arg ListJobToFacebookAdByJobIDParams) ([]JobToFacebookAd, error) {
	rows, err := q.db.QueryContext(ctx, listJobToFacebookAdByJobID, arg.JobID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []JobToFacebookAd{}
	for rows.Next() {
		var i JobToFacebookAd
		if err := rows.Scan(
			&i.ID,
			&i.JobID,
			&i.AdID,
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

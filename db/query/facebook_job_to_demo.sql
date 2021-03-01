-- name: CreateJobToFacebookDemo :one
INSERT INTO "JobToFacebookDemo" (
  job_id,
  ad_demo_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetJobToFacebookDemo :one
SELECT * FROM "JobToFacebookDemo"
WHERE id = $1 LIMIT 1;


-- name: ListJobToFacebookDemo :many
SELECT * FROM "JobToFacebookDemo"
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: ListJobToFacebookDemoByJobID :many
SELECT * FROM "JobToFacebookDemo"
WHERE job_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: DeleteJobToFacebookDemo :exec
DELETE FROM "JobToFacebookDemo"
WHERE id = $1;
-- name: CreateJobToFacebookAd :one
INSERT INTO "JobToFacebookAd" (
  job_id,
  ad_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetJobToFacebookAd :one
SELECT * FROM "JobToFacebookAd"
WHERE id = $1 LIMIT 1;


-- name: ListJobToFacebookAd :many
SELECT * FROM "JobToFacebookAd"
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: ListJobToFacebookAdByJobID :many
SELECT * FROM "JobToFacebookAd"
WHERE job_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: DeleteJobToFacebookAd :exec
DELETE FROM "JobToFacebookAd"
WHERE id = $1;
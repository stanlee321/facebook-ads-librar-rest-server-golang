-- name: CreateJobToFacebookRegion :one
INSERT INTO "JobToFacebookRe" (
  job_id,
  ad_region_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetJobToFacebookRegion :one
SELECT * FROM "JobToFacebookRe"
WHERE id = $1 LIMIT 1;


-- name: ListJobToFacebookRegion :many
SELECT * FROM "JobToFacebookRe"
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: ListJobToFacebookRegionByJobID :many
SELECT * FROM "JobToFacebookRe"
WHERE job_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: DeleteJobToFacebookRegion :exec
DELETE FROM "JobToFacebookRe"
WHERE id = $1;
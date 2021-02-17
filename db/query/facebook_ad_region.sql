-- name: CreateFacebookRegion :one
INSERT INTO "FacebookRegions" (
  ad_id,
  page_id,
  region,
  percentage,
  ad_delivery_start_time
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetFacebookRegion :one
SELECT * FROM "FacebookRegions"
WHERE id = $1 LIMIT 1;

-- name: ListFacebookRegions :many
SELECT * FROM "FacebookRegions"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListFacebookRegionsByAdID :many
SELECT * FROM "FacebookRegions"
WHERE ad_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListFacebookRegionsByPageID :many
SELECT * FROM "FacebookRegions"
WHERE page_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteFaceookRegion :exec
DELETE FROM "FacebookDemos"
WHERE id = $1;
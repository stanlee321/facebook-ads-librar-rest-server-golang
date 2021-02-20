-- name: CreateFacebookJob :one
INSERT INTO "FacebookJob" (
  search_terms,
  access_token,
  page_total,
  search_total,
  ad_active_status,
  ad_delivery_date_max,
  ad_delivery_date_min,
  ad_reached_countries
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetFacebookJob :one
SELECT * FROM "FacebookJob"
WHERE id = $1 LIMIT 1;

-- name: GetPastFacebookJob :one
SELECT * FROM "FacebookJob"
WHERE search_terms = $1 AND
  page_total = $2 AND
  search_total = $3 AND
  ad_active_status = $4 AND
  ad_delivery_date_max = $5 AND
  ad_delivery_date_min = $6 AND
  ad_reached_countries = $7
LIMIT 1;



-- name: ListFacebookJobs :many
SELECT * FROM "FacebookJob"
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: ListFacebookJobsBySearch :many
SELECT * FROM "FacebookJob"
WHERE search_terms = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListFacebookJobsByToken :many
SELECT * FROM "FacebookJob"
WHERE access_token = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: DeleteFaceookJob :exec
DELETE FROM "FacebookJob"
WHERE id = $1;
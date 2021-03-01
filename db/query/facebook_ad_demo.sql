-- name: CreateFacebookDemo :one
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
RETURNING *;

-- name: GetFacebookDemo :one
SELECT * FROM "FacebookDemos"
WHERE id = $1 LIMIT 1;


-- name: ListFacebookDemos :many
SELECT * FROM "FacebookDemos"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListFacebookDemosByAdID :many
SELECT * FROM "FacebookDemos"
WHERE ad_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;



-- name: ListFacebookDemosByPageID :many
SELECT * FROM "FacebookDemos"
WHERE page_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteFaceookDemo :exec
DELETE FROM "FacebookDemos"
WHERE id = $1;
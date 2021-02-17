-- name: CreateFacebookAd :one
INSERT INTO "FacebookAd" (
  ad_id,
  page_id,
  page_name,
  ad_snapshot_url,
  ad_creative_body,
  ad_creative_link_caption,
  ad_creative_link_description,
  ad_creative_link_title,
  ad_delivery_start_time,
  ad_delivery_stop_time,
  funding_entity,
  impressions_min,
  spend_min,
  spend_max,
  currency,
  ad_url,
  social_media_facebook,
  social_media_instagram,
  social_media_whatsapp,
  search_terms
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20
)
RETURNING *;

-- name: GetFacebookAd :one
SELECT * FROM "FacebookAd"
WHERE id = $1 LIMIT 1;


-- name: ListFacebookAds :many
SELECT * FROM "FacebookAd"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListFacebookAdsByAdID :many
SELECT * FROM "FacebookAd"
WHERE ad_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListFacebookAdsByPageID :many
SELECT * FROM "FacebookAd"
WHERE page_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListFacebookAdsByPageName :many
SELECT * FROM "FacebookAd"
WHERE page_name = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteFaceookAd :exec
DELETE FROM "FacebookAd"
WHERE id = $1;
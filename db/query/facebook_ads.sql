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
  impressions_max,
  spend_min,
  spend_max,
  currency,
  ad_url,
  social_media_facebook,
  social_media_instagram,
  social_media_whatsapp,
  search_terms,
  ad_creation_time,
  potential_reach_max,
  potential_reach_min
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24
)
RETURNING *;

-- name: GetFacebookAd :one
SELECT * FROM "FacebookAd"
WHERE ad_id = $1 LIMIT 1;


-- name: ListFacebookAds :many
SELECT * FROM "FacebookAd"
LIMIT $1
OFFSET $2;



-- name: ListFacebookAdsByPageID :many
SELECT * FROM "FacebookAd"
WHERE page_id = $1
LIMIT $2
OFFSET $3;

-- name: ListFacebookAdsByPageName :many
SELECT * FROM "FacebookAd"
WHERE page_name = $1
LIMIT $2
OFFSET $3;

-- name: DeleteFaceookAd :exec
DELETE FROM "FacebookAd"
WHERE ad_id = $1;
// Code generated by sqlc. DO NOT EDIT.
// source: facebook_ads.sql

package db

import (
	"context"
	"database/sql"
)

const createFacebookAd = `-- name: CreateFacebookAd :one
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
RETURNING ad_id, page_id, page_name, ad_snapshot_url, ad_creative_body, ad_creative_link_caption, ad_creative_link_description, ad_creative_link_title, ad_delivery_start_time, ad_delivery_stop_time, funding_entity, impressions_min, impressions_max, spend_min, spend_max, currency, ad_url, social_media_facebook, social_media_instagram, social_media_whatsapp, search_terms, ad_creation_time, potential_reach_max, potential_reach_min, created_at
`

type CreateFacebookAdParams struct {
	AdID                      int64          `json:"ad_id"`
	PageID                    sql.NullInt64  `json:"page_id"`
	PageName                  sql.NullString `json:"page_name"`
	AdSnapshotUrl             sql.NullString `json:"ad_snapshot_url"`
	AdCreativeBody            sql.NullString `json:"ad_creative_body"`
	AdCreativeLinkCaption     sql.NullString `json:"ad_creative_link_caption"`
	AdCreativeLinkDescription sql.NullString `json:"ad_creative_link_description"`
	AdCreativeLinkTitle       sql.NullString `json:"ad_creative_link_title"`
	AdDeliveryStartTime       sql.NullString `json:"ad_delivery_start_time"`
	AdDeliveryStopTime        sql.NullString `json:"ad_delivery_stop_time"`
	FundingEntity             sql.NullString `json:"funding_entity"`
	ImpressionsMin            sql.NullInt32  `json:"impressions_min"`
	ImpressionsMax            sql.NullInt32  `json:"impressions_max"`
	SpendMin                  sql.NullInt32  `json:"spend_min"`
	SpendMax                  sql.NullInt32  `json:"spend_max"`
	Currency                  sql.NullString `json:"currency"`
	AdUrl                     sql.NullString `json:"ad_url"`
	SocialMediaFacebook       sql.NullString `json:"social_media_facebook"`
	SocialMediaInstagram      sql.NullString `json:"social_media_instagram"`
	SocialMediaWhatsapp       sql.NullString `json:"social_media_whatsapp"`
	SearchTerms               sql.NullString `json:"search_terms"`
	AdCreationTime            sql.NullString `json:"ad_creation_time"`
	PotentialReachMax         sql.NullInt32  `json:"potential_reach_max"`
	PotentialReachMin         sql.NullInt32  `json:"potential_reach_min"`
}

func (q *Queries) CreateFacebookAd(ctx context.Context, arg CreateFacebookAdParams) (FacebookAd, error) {
	row := q.db.QueryRowContext(ctx, createFacebookAd,
		arg.AdID,
		arg.PageID,
		arg.PageName,
		arg.AdSnapshotUrl,
		arg.AdCreativeBody,
		arg.AdCreativeLinkCaption,
		arg.AdCreativeLinkDescription,
		arg.AdCreativeLinkTitle,
		arg.AdDeliveryStartTime,
		arg.AdDeliveryStopTime,
		arg.FundingEntity,
		arg.ImpressionsMin,
		arg.ImpressionsMax,
		arg.SpendMin,
		arg.SpendMax,
		arg.Currency,
		arg.AdUrl,
		arg.SocialMediaFacebook,
		arg.SocialMediaInstagram,
		arg.SocialMediaWhatsapp,
		arg.SearchTerms,
		arg.AdCreationTime,
		arg.PotentialReachMax,
		arg.PotentialReachMin,
	)
	var i FacebookAd
	err := row.Scan(
		&i.AdID,
		&i.PageID,
		&i.PageName,
		&i.AdSnapshotUrl,
		&i.AdCreativeBody,
		&i.AdCreativeLinkCaption,
		&i.AdCreativeLinkDescription,
		&i.AdCreativeLinkTitle,
		&i.AdDeliveryStartTime,
		&i.AdDeliveryStopTime,
		&i.FundingEntity,
		&i.ImpressionsMin,
		&i.ImpressionsMax,
		&i.SpendMin,
		&i.SpendMax,
		&i.Currency,
		&i.AdUrl,
		&i.SocialMediaFacebook,
		&i.SocialMediaInstagram,
		&i.SocialMediaWhatsapp,
		&i.SearchTerms,
		&i.AdCreationTime,
		&i.PotentialReachMax,
		&i.PotentialReachMin,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFaceookAd = `-- name: DeleteFaceookAd :exec
DELETE FROM "FacebookAd"
WHERE ad_id = $1
`

func (q *Queries) DeleteFaceookAd(ctx context.Context, adID int64) error {
	_, err := q.db.ExecContext(ctx, deleteFaceookAd, adID)
	return err
}

const getFacebookAd = `-- name: GetFacebookAd :one
SELECT ad_id, page_id, page_name, ad_snapshot_url, ad_creative_body, ad_creative_link_caption, ad_creative_link_description, ad_creative_link_title, ad_delivery_start_time, ad_delivery_stop_time, funding_entity, impressions_min, impressions_max, spend_min, spend_max, currency, ad_url, social_media_facebook, social_media_instagram, social_media_whatsapp, search_terms, ad_creation_time, potential_reach_max, potential_reach_min, created_at FROM "FacebookAd"
WHERE ad_id = $1 LIMIT 1
`

func (q *Queries) GetFacebookAd(ctx context.Context, adID int64) (FacebookAd, error) {
	row := q.db.QueryRowContext(ctx, getFacebookAd, adID)
	var i FacebookAd
	err := row.Scan(
		&i.AdID,
		&i.PageID,
		&i.PageName,
		&i.AdSnapshotUrl,
		&i.AdCreativeBody,
		&i.AdCreativeLinkCaption,
		&i.AdCreativeLinkDescription,
		&i.AdCreativeLinkTitle,
		&i.AdDeliveryStartTime,
		&i.AdDeliveryStopTime,
		&i.FundingEntity,
		&i.ImpressionsMin,
		&i.ImpressionsMax,
		&i.SpendMin,
		&i.SpendMax,
		&i.Currency,
		&i.AdUrl,
		&i.SocialMediaFacebook,
		&i.SocialMediaInstagram,
		&i.SocialMediaWhatsapp,
		&i.SearchTerms,
		&i.AdCreationTime,
		&i.PotentialReachMax,
		&i.PotentialReachMin,
		&i.CreatedAt,
	)
	return i, err
}

const listFacebookAds = `-- name: ListFacebookAds :many
SELECT ad_id, page_id, page_name, ad_snapshot_url, ad_creative_body, ad_creative_link_caption, ad_creative_link_description, ad_creative_link_title, ad_delivery_start_time, ad_delivery_stop_time, funding_entity, impressions_min, impressions_max, spend_min, spend_max, currency, ad_url, social_media_facebook, social_media_instagram, social_media_whatsapp, search_terms, ad_creation_time, potential_reach_max, potential_reach_min, created_at FROM "FacebookAd"
LIMIT $1
OFFSET $2
`

type ListFacebookAdsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFacebookAds(ctx context.Context, arg ListFacebookAdsParams) ([]FacebookAd, error) {
	rows, err := q.db.QueryContext(ctx, listFacebookAds, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FacebookAd{}
	for rows.Next() {
		var i FacebookAd
		if err := rows.Scan(
			&i.AdID,
			&i.PageID,
			&i.PageName,
			&i.AdSnapshotUrl,
			&i.AdCreativeBody,
			&i.AdCreativeLinkCaption,
			&i.AdCreativeLinkDescription,
			&i.AdCreativeLinkTitle,
			&i.AdDeliveryStartTime,
			&i.AdDeliveryStopTime,
			&i.FundingEntity,
			&i.ImpressionsMin,
			&i.ImpressionsMax,
			&i.SpendMin,
			&i.SpendMax,
			&i.Currency,
			&i.AdUrl,
			&i.SocialMediaFacebook,
			&i.SocialMediaInstagram,
			&i.SocialMediaWhatsapp,
			&i.SearchTerms,
			&i.AdCreationTime,
			&i.PotentialReachMax,
			&i.PotentialReachMin,
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

const listFacebookAdsByAdID = `-- name: ListFacebookAdsByAdID :many
SELECT ad_id, page_id, page_name, ad_snapshot_url, ad_creative_body, ad_creative_link_caption, ad_creative_link_description, ad_creative_link_title, ad_delivery_start_time, ad_delivery_stop_time, funding_entity, impressions_min, impressions_max, spend_min, spend_max, currency, ad_url, social_media_facebook, social_media_instagram, social_media_whatsapp, search_terms, ad_creation_time, potential_reach_max, potential_reach_min, created_at FROM "FacebookAd"
WHERE ad_id = $1
LIMIT $2
OFFSET $3
`

type ListFacebookAdsByAdIDParams struct {
	AdID   int64 `json:"ad_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFacebookAdsByAdID(ctx context.Context, arg ListFacebookAdsByAdIDParams) ([]FacebookAd, error) {
	rows, err := q.db.QueryContext(ctx, listFacebookAdsByAdID, arg.AdID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FacebookAd{}
	for rows.Next() {
		var i FacebookAd
		if err := rows.Scan(
			&i.AdID,
			&i.PageID,
			&i.PageName,
			&i.AdSnapshotUrl,
			&i.AdCreativeBody,
			&i.AdCreativeLinkCaption,
			&i.AdCreativeLinkDescription,
			&i.AdCreativeLinkTitle,
			&i.AdDeliveryStartTime,
			&i.AdDeliveryStopTime,
			&i.FundingEntity,
			&i.ImpressionsMin,
			&i.ImpressionsMax,
			&i.SpendMin,
			&i.SpendMax,
			&i.Currency,
			&i.AdUrl,
			&i.SocialMediaFacebook,
			&i.SocialMediaInstagram,
			&i.SocialMediaWhatsapp,
			&i.SearchTerms,
			&i.AdCreationTime,
			&i.PotentialReachMax,
			&i.PotentialReachMin,
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

const listFacebookAdsByPageID = `-- name: ListFacebookAdsByPageID :many
SELECT ad_id, page_id, page_name, ad_snapshot_url, ad_creative_body, ad_creative_link_caption, ad_creative_link_description, ad_creative_link_title, ad_delivery_start_time, ad_delivery_stop_time, funding_entity, impressions_min, impressions_max, spend_min, spend_max, currency, ad_url, social_media_facebook, social_media_instagram, social_media_whatsapp, search_terms, ad_creation_time, potential_reach_max, potential_reach_min, created_at FROM "FacebookAd"
WHERE page_id = $1
LIMIT $2
OFFSET $3
`

type ListFacebookAdsByPageIDParams struct {
	PageID sql.NullInt64 `json:"page_id"`
	Limit  int32         `json:"limit"`
	Offset int32         `json:"offset"`
}

func (q *Queries) ListFacebookAdsByPageID(ctx context.Context, arg ListFacebookAdsByPageIDParams) ([]FacebookAd, error) {
	rows, err := q.db.QueryContext(ctx, listFacebookAdsByPageID, arg.PageID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FacebookAd{}
	for rows.Next() {
		var i FacebookAd
		if err := rows.Scan(
			&i.AdID,
			&i.PageID,
			&i.PageName,
			&i.AdSnapshotUrl,
			&i.AdCreativeBody,
			&i.AdCreativeLinkCaption,
			&i.AdCreativeLinkDescription,
			&i.AdCreativeLinkTitle,
			&i.AdDeliveryStartTime,
			&i.AdDeliveryStopTime,
			&i.FundingEntity,
			&i.ImpressionsMin,
			&i.ImpressionsMax,
			&i.SpendMin,
			&i.SpendMax,
			&i.Currency,
			&i.AdUrl,
			&i.SocialMediaFacebook,
			&i.SocialMediaInstagram,
			&i.SocialMediaWhatsapp,
			&i.SearchTerms,
			&i.AdCreationTime,
			&i.PotentialReachMax,
			&i.PotentialReachMin,
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

const listFacebookAdsByPageName = `-- name: ListFacebookAdsByPageName :many
SELECT ad_id, page_id, page_name, ad_snapshot_url, ad_creative_body, ad_creative_link_caption, ad_creative_link_description, ad_creative_link_title, ad_delivery_start_time, ad_delivery_stop_time, funding_entity, impressions_min, impressions_max, spend_min, spend_max, currency, ad_url, social_media_facebook, social_media_instagram, social_media_whatsapp, search_terms, ad_creation_time, potential_reach_max, potential_reach_min, created_at FROM "FacebookAd"
WHERE page_name = $1
LIMIT $2
OFFSET $3
`

type ListFacebookAdsByPageNameParams struct {
	PageName sql.NullString `json:"page_name"`
	Limit    int32          `json:"limit"`
	Offset   int32          `json:"offset"`
}

func (q *Queries) ListFacebookAdsByPageName(ctx context.Context, arg ListFacebookAdsByPageNameParams) ([]FacebookAd, error) {
	rows, err := q.db.QueryContext(ctx, listFacebookAdsByPageName, arg.PageName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FacebookAd{}
	for rows.Next() {
		var i FacebookAd
		if err := rows.Scan(
			&i.AdID,
			&i.PageID,
			&i.PageName,
			&i.AdSnapshotUrl,
			&i.AdCreativeBody,
			&i.AdCreativeLinkCaption,
			&i.AdCreativeLinkDescription,
			&i.AdCreativeLinkTitle,
			&i.AdDeliveryStartTime,
			&i.AdDeliveryStopTime,
			&i.FundingEntity,
			&i.ImpressionsMin,
			&i.ImpressionsMax,
			&i.SpendMin,
			&i.SpendMax,
			&i.Currency,
			&i.AdUrl,
			&i.SocialMediaFacebook,
			&i.SocialMediaInstagram,
			&i.SocialMediaWhatsapp,
			&i.SearchTerms,
			&i.AdCreationTime,
			&i.PotentialReachMax,
			&i.PotentialReachMin,
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

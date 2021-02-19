package api

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
	pb "github.com/stanlee321/facebook-ads-server/pkg/api/v1"

	"github.com/gin-gonic/gin"
)

type createFacebookAdRequest struct {
	AdID                      sql.NullInt64  `json:"ad_id" binding:"required"`
	PageID                    sql.NullInt64  `json:"page_id" binding:"required"`
	PageName                  sql.NullString `json:"page_name" binding:"required"`
	AdSnapshotURL             sql.NullString `json:"ad_snapshot_url"`
	AdCreativeBody            sql.NullString `json:"ad_creative_body"`
	AdCreativeLinkCaption     sql.NullString `json:"ad_creative_link_caption"`
	AdCreativeLinkDescription sql.NullString `json:"ad_creative_link_description"`
	AdCreativeLinkTitle       sql.NullString `json:"ad_creative_link_title"`
	AdDeliveryStartTime       sql.NullString `json:"ad_delivery_start_time"`
	AdDeliveryStopTime        sql.NullString `json:"ad_delivery_stop_time"`
	FundingEntity             sql.NullString `json:"funding_entity"`
	ImpressionsMin            sql.NullString `json:"impressions_min"`
	SpendMin                  sql.NullInt64  `json:"spend_min"`
	SpendMax                  sql.NullInt64  `json:"spend_max"`
	Currency                  sql.NullString `json:"currency"`
	AdURL                     sql.NullString `json:"ad_url"`
	SocialMediaFacebook       sql.NullString `json:"social_media_facebook"`
	SocialMediaInstagram      sql.NullString `json:"social_media_instagram"`
	SocialMediaWhatsapp       sql.NullString `json:"social_media_whatsapp"`
	SearchTerms               sql.NullString `json:"search_terms"`
}

func (server *Server) createFacebookAd(ctx *gin.Context) {
	var req createFacebookAdRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateFacebookAdParams{
		AdID:                      req.AdID,
		PageID:                    req.PageID,
		PageName:                  req.PageName,
		AdSnapshotUrl:             req.AdSnapshotURL,
		AdCreativeBody:            req.AdCreativeBody,
		AdCreativeLinkCaption:     req.AdCreativeLinkCaption,
		AdCreativeLinkDescription: req.AdCreativeLinkDescription,
		AdCreativeLinkTitle:       req.AdCreativeLinkTitle,
		AdDeliveryStartTime:       req.AdDeliveryStartTime,
		AdDeliveryStopTime:        req.AdDeliveryStopTime,
		FundingEntity:             req.FundingEntity,
		ImpressionsMin:            req.ImpressionsMin,
		SpendMin:                  req.SpendMin,
		SpendMax:                  req.SpendMax,
		Currency:                  req.Currency,
		AdUrl:                     req.AdURL,
		SocialMediaFacebook:       req.SocialMediaFacebook,
		SocialMediaInstagram:      req.SocialMediaInstagram,
		SocialMediaWhatsapp:       req.SocialMediaWhatsapp,
		SearchTerms:               req.SearchTerms,
	}

	ad, err := server.store.CreateFacebookAd(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ad)
}

type getFacebookAdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFacebookAd(ctx *gin.Context) {
	var req getFacebookAdRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ad, err := server.store.GetFacebookAd(ctx, req.ID)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ad)
}

type listFacebookAdByPageIDRequest struct {
	PageLocation int32 `form:"page_location" binding:"required,min=1"`
	PageSize     int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFacebookAds(ctx *gin.Context) {
	var req listFacebookAdByPageIDRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListFacebookAdsParams{
		Limit:  req.PageSize,
		Offset: (req.PageLocation - 1) * req.PageSize,
	}

	ads, err := server.store.ListFacebookAds(ctx, args)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ads)
}

type listFacebookAdsByPageIDRequest struct {
	PageID       int64 `form:"page_id" binding:"required"`
	PageLocation int32 `form:"page_location" binding:"required,min=1"`
	PageSize     int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFacebookAdsByPageID(ctx *gin.Context) {
	var req listFacebookAdsByPageIDRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListFacebookAdsByPageIDParams{
		PageID: sql.NullInt64{Int64: req.PageID, Valid: true},
		Limit:  req.PageSize,
		Offset: (req.PageLocation - 1) * req.PageSize,
	}

	ads, err := server.store.ListFacebookAdsByPageID(ctx, args)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ads)
}

type listFacebookAdsByPageNameRequest struct {
	PageName     string `form:"page_name" binding:"required"`
	PageLocation int32  `form:"page_location" binding:"required,min=1"`
	PageSize     int32  `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFacebookAdsByPageName(ctx *gin.Context) {
	var req listFacebookAdsByPageNameRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListFacebookAdsByPageNameParams{
		PageName: sql.NullString{String: req.PageName, Valid: true},
		Limit:    req.PageSize,
		Offset:   (req.PageLocation - 1) * req.PageSize,
	}

	ads, err := server.store.ListFacebookAdsByPageName(ctx, args)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ads)

}

type listFacebookAdsByAdIDRequest struct {
	AdID         int64 `form:"ad_id" binding:"required"`
	PageLocation int32 `form:"page_location" binding:"required,min=1"`
	PageSize     int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFacebookAdsByAdID(ctx *gin.Context) {
	var req listFacebookAdsByAdIDRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListFacebookAdsByAdIDParams{
		AdID:   sql.NullInt64{Int64: req.AdID, Valid: true},
		Limit:  req.PageSize,
		Offset: (req.PageLocation - 1) * req.PageSize,
	}

	ads, err := server.store.ListFacebookAdsByAdID(ctx, args)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ads)
}

type deleteFacebookAdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteFacebookAd(ctx *gin.Context) {
	var req deleteFacebookAdRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteFaceookAd(ctx, req.ID)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}

type getSearchRequest struct {
	SearchTerms        string `json:"search_terms" binding:"required"`
	AccessToken        string `json:"access_token" binding:"required"`
	PageTotal          int32  `json:"page_total" binding:"required,min=1,max=1000"`
	SearchTotal        int32  `json:"search_total" binding:"required,min=1,max=5000"`
	AdActiveStatus     string `json:"ad_active_status" binding:"required,oneof=ACTIVE INACTIVE ALL"`
	AdDeliveryDateMax  string `json:"ad_delivery_date_max" binding:"required"`
	AdDeliveryDateMin  string `json:"ad_delivery_date_min" binding:"required"`
	AdReachedCountries string `json:"ad_reached_countries" binding:"required,oneof=BO MX"`
}

func (server *Server) getSearch(ctx *gin.Context) {
	var req getSearchRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	adStream, err := server.facebookClient.CreateFacebookAd(ctx, &pb.CreateFacebookAdRequest{
		SearchTerms:        req.SearchTerms,
		AccessToken:        req.AccessToken,
		PageTotal:          req.PageTotal,
		SearchTotal:        req.SearchTotal,
		AdActiveStatus:     req.AdActiveStatus,
		AdDeliveryDateMax:  req.AdDeliveryDateMax,
		AdDeliveryDateMin:  req.AdDeliveryDateMin,
		AdReachedCountries: req.AdReachedCountries,
	})

	if err != nil {
		log.Print("ERROR FROM GRPC CALL", err)

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var adResults []*pb.FacebookAd

	for {
		ad, err := adStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Print("ERROR FROM GRPC CALL", err)

			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// fmt.Print(ad)

		adResults = append(adResults, ad.FacebookAd)
	}

	ctx.JSON(http.StatusOK, adResults)
}

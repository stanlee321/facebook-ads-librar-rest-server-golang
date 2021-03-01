package api

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
	pb "github.com/stanlee321/facebook-ads-server/pkg/ads/api/v1"
)

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

	ctx.JSON(http.StatusOK, convertFBAdDBtoFBAdPB(ad))
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
	ctx.JSON(http.StatusOK, convertFBAdDBtoFBAdPBList(ads))
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

	ctx.JSON(http.StatusOK, convertFBAdDBtoFBAdPBList(ads))
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

	ctx.JSON(http.StatusOK, convertFBAdDBtoFBAdPBList(ads))
}

type listFacebookJobsRequest struct {
	PageLocation int32 `form:"page_location" binding:"required,min=1"`
	PageSize     int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFacebookJobs(ctx *gin.Context) {
	var req listFacebookJobsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListFacebookJobsParams{
		Limit:  req.PageSize,
		Offset: (req.PageLocation - 1) * req.PageSize,
	}

	jobsDB, err := server.store.ListFacebookJobs(ctx, args)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, convertJobDBtoJobResponse(jobsDB))

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

type createJobRequest struct {
	SearchTerms        string `json:"search_terms" binding:"required"`
	AccessToken        string `json:"access_token" binding:"required"`
	PageTotal          int32  `json:"page_total" binding:"required,min=1,max=1000"`
	SearchTotal        int32  `json:"search_total" binding:"required,min=1,max=5000"`
	AdActiveStatus     string `json:"ad_active_status" binding:"required,oneof=ACTIVE INACTIVE ALL"`
	AdDeliveryDateMax  string `json:"ad_delivery_date_max" binding:"required"`
	AdDeliveryDateMin  string `json:"ad_delivery_date_min" binding:"required"`
	AdReachedCountries string `json:"ad_reached_countries" binding:"required,oneof=BO MX"`
}

type createJobResponse struct {
	SearchTerms string `json:"search_terms"`
	JobID       int64  `json:"job_id"`
	AccessToken string `json:"access_token"`
	TotalAds    int    `json:"total_ads"`
}

func (server *Server) createJob(ctx *gin.Context) {
	var req createJobRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check if job exists

	argsJobQuery := db.GetPastFacebookJobParams{
		SearchTerms:        sql.NullString{String: req.SearchTerms, Valid: true},
		PageTotal:          sql.NullInt64{Int64: int64(req.PageTotal), Valid: true},
		SearchTotal:        sql.NullInt64{Int64: int64(req.SearchTotal), Valid: true},
		AdActiveStatus:     sql.NullString{String: req.AdActiveStatus, Valid: true},
		AdDeliveryDateMax:  sql.NullString{String: req.AdDeliveryDateMax, Valid: true},
		AdDeliveryDateMin:  sql.NullString{String: req.AdDeliveryDateMin, Valid: true},
		AdReachedCountries: sql.NullString{String: req.AdReachedCountries, Valid: true},
	}

	jobDB, err := server.store.GetPastFacebookJob(ctx, argsJobQuery)

	if err == sql.ErrNoRows {

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

		for {

			adSet, err := adStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Print("ERROR FROM GRPC call stream...", err)

				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(err))
					return
				}
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			// Create Job Args
			argsJobCreation := db.CreateFacebookJobParams{
				SearchTerms:        sql.NullString{String: req.SearchTerms, Valid: true},
				AccessToken:        sql.NullString{String: req.AccessToken, Valid: true},
				PageTotal:          sql.NullInt64{Int64: int64(req.PageTotal), Valid: true},
				SearchTotal:        sql.NullInt64{Int64: int64(req.SearchTotal), Valid: true},
				AdActiveStatus:     sql.NullString{String: req.AdActiveStatus, Valid: true},
				AdDeliveryDateMax:  sql.NullString{String: req.AdDeliveryDateMax, Valid: true},
				AdDeliveryDateMin:  sql.NullString{String: req.AdDeliveryDateMin, Valid: true},
				AdReachedCountries: sql.NullString{String: req.AdReachedCountries, Valid: true},
				TotalFoundAds:      sql.NullInt64{Int64: int64(len(adSet.FacebookAd)), Valid: true},
			}

			// Create Job in DB
			newJob, err := server.store.CreateFacebookJob(ctx, argsJobCreation)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			// For Ads
			for _, facebookAd := range adSet.FacebookAd {

				adID, err := stringToBigInt(facebookAd.AdId)

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				pageID, err := stringToBigInt(facebookAd.PageId)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				// Create Ad in DB
				args := db.CreateFacebookAdParams{
					AdID:                      adID,
					PageID:                    sql.NullInt64{Int64: pageID, Valid: true},
					PageName:                  sql.NullString{String: facebookAd.PageName, Valid: true},
					AdSnapshotUrl:             sql.NullString{String: facebookAd.AdSnapshotUrl, Valid: true},
					AdCreativeBody:            sql.NullString{String: facebookAd.AdCreativeBody, Valid: true},
					AdCreativeLinkCaption:     sql.NullString{String: facebookAd.AdCreativeLinkCaption, Valid: true},
					AdCreativeLinkDescription: sql.NullString{String: facebookAd.AdCreativeLinkDescription, Valid: true},
					AdCreativeLinkTitle:       sql.NullString{String: facebookAd.AdCreativeLinkTitle, Valid: true},
					AdDeliveryStartTime:       sql.NullString{String: facebookAd.AdDeliveryStartTime, Valid: true},
					AdDeliveryStopTime:        sql.NullString{String: facebookAd.AdDeliveryStopTime, Valid: true},
					FundingEntity:             sql.NullString{String: facebookAd.FundingEntity, Valid: true},
					ImpressionsMax:            sql.NullInt32{Int32: facebookAd.ImpressionsMax, Valid: true},
					ImpressionsMin:            sql.NullInt32{Int32: facebookAd.ImpressionsMin, Valid: true},
					SpendMin:                  sql.NullInt32{Int32: facebookAd.SpendMin, Valid: true},
					SpendMax:                  sql.NullInt32{Int32: facebookAd.SpendMax, Valid: true},
					Currency:                  sql.NullString{String: facebookAd.Currency, Valid: true},
					AdUrl:                     sql.NullString{String: facebookAd.AdUrl, Valid: true},
					SocialMediaFacebook:       sql.NullString{String: facebookAd.SocialMediaFacebook, Valid: true},
					SocialMediaInstagram:      sql.NullString{String: facebookAd.SocialMediaInstagram, Valid: true},
					SocialMediaWhatsapp:       sql.NullString{String: facebookAd.SocialMediaWhatsapp, Valid: true},
					SearchTerms:               sql.NullString{String: facebookAd.SearchTerms, Valid: true},
					AdCreationTime:            sql.NullString{String: facebookAd.AdCreationTime, Valid: true},
					PotentialReachMax:         sql.NullInt32{Int32: facebookAd.PotentialReachMax, Valid: true},
					PotentialReachMin:         sql.NullInt32{Int32: facebookAd.PotentialReachMin, Valid: true},
				}

				_, err = server.store.CreateFacebookAd(ctx, args)

				if err == nil {
					// Create Job To Facebook ad Map
					argsJTF := db.CreateJobToFacebookAdParams{
						JobID: sql.NullInt64{Int64: newJob.ID, Valid: true},
						AdID:  sql.NullInt64{Int64: adID, Valid: true},
					}

					// Save in DB
					_, err = server.store.CreateJobToFacebookAd(ctx, argsJTF)

					// If error in save to db
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, errorResponse(err))
						return
					}

				} else {
					if strings.Contains("pq: duplicate key value violates unique ", err.Error()) {
						// Create Job To Facebook ad Map
						argsJTF := db.CreateJobToFacebookAdParams{
							JobID: sql.NullInt64{Int64: newJob.ID, Valid: true},
							AdID:  sql.NullInt64{Int64: adID, Valid: true},
						}

						// Save in DB
						_, err = server.store.CreateJobToFacebookAd(ctx, argsJTF)

						// If error in save to db
						if err != nil {
							ctx.JSON(http.StatusInternalServerError, errorResponse(err))
							return
						}
					} else {
						// Another kind of error
						ctx.JSON(http.StatusInternalServerError, errorResponse(err))
						return
					}

				}

			}

			// Create Demo inDB
			for _, facebookDemo := range adSet.FacebookAdDemo {

				adID, err := stringToBigInt(facebookDemo.AdId)

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				pageID, err := stringToBigInt(facebookDemo.PageId)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				argsDemo := db.CreateFacebookDemoParams{
					AdID:                sql.NullInt64{Int64: adID, Valid: true},
					PageID:              sql.NullInt64{Int64: pageID, Valid: true},
					Age:                 sql.NullString{String: facebookDemo.Age, Valid: true},
					Gender:              sql.NullString{String: facebookDemo.Gender, Valid: true},
					Percentage:          sql.NullString{String: fmt.Sprint(facebookDemo.Percentage), Valid: true},
					AdDeliveryStartTime: sql.NullString{String: facebookDemo.AdDeliveryStartTime, Valid: true},
				}

				demoResponse, err := server.store.CreateFacebookDemo(ctx, argsDemo)

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				// Create Job To Facebook Demo ad Map
				argsJobToDemo := db.CreateJobToFacebookDemoParams{
					JobID:    sql.NullInt64{Int64: newJob.ID, Valid: true},
					AdDemoID: sql.NullInt64{Int64: demoResponse.ID, Valid: true},
				}

				// Save in DB
				_, err = server.store.CreateJobToFacebookDemo(ctx, argsJobToDemo)

				// If error in save to db
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

			}

			// Create Region inDB
			for _, facebookRegion := range adSet.FacebookAdRegion {
				adID, err := stringToBigInt(facebookRegion.AdId)

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				pageID, err := stringToBigInt(facebookRegion.PageId)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				argsregion := db.CreateFacebookRegionParams{
					AdID:                sql.NullInt64{Int64: adID, Valid: true},
					PageID:              sql.NullInt64{Int64: pageID, Valid: true},
					Region:              sql.NullString{String: facebookRegion.Region, Valid: true},
					Percentage:          sql.NullString{String: fmt.Sprint(facebookRegion.Percentage), Valid: true},
					AdDeliveryStartTime: sql.NullString{String: facebookRegion.AdDeliveryStartTime, Valid: true},
				}
				regionResponse, err := server.store.CreateFacebookRegion(ctx, argsregion)

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				// Create Job To Facebook Region ad Map
				argsJobToRegion := db.CreateJobToFacebookRegionParams{
					JobID:      sql.NullInt64{Int64: newJob.ID, Valid: true},
					AdRegionID: sql.NullInt64{Int64: regionResponse.ID, Valid: true},
				}

				// Save in DB
				_, err = server.store.CreateJobToFacebookRegion(ctx, argsJobToRegion)

				// If error in save to db
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

			}

			jobResponse := createJobResponse{
				SearchTerms: newJob.SearchTerms.String,
				JobID:       newJob.ID,
				AccessToken: newJob.AccessToken.String,
				TotalAds:    len(adSet.FacebookAd),
			}

			ctx.JSON(http.StatusOK, jobResponse)
			return
		}
	}
	argsListJobs := db.ListJobToFacebookAdByJobIDParams{
		JobID:  sql.NullInt64{Int64: jobDB.ID, Valid: true},
		Limit:  99999999, // TODO CHECK IF FILTER FOR GET COUNT WITHOUT QUERY
		Offset: 1,        // TODO CHECK IF FILTER FOR GET COUNT WITHOUT QUERY
	}
	// If search does not exists
	fbAds, err := server.store.ListJobToFacebookAdByJobID(ctx, argsListJobs)

	var response createJobResponse

	response.AccessToken = req.AccessToken
	response.JobID = jobDB.ID
	response.SearchTerms = req.SearchTerms
	response.TotalAds = len(fbAds)

	ctx.JSON(http.StatusOK, response)
}

func stringToBigInt(stringInput string) (int64, error) {

	i, err := strconv.ParseInt(stringInput, 10, 64)

	if err != nil {
		return 0, err
	}

	return i, nil
}

func convertFBAdDBtoFBAdPBList(ads []db.FacebookAd) []*pb.FacebookAd {

	var adsPB []*pb.FacebookAd

	for _, ad := range ads {

		newPB := convertFBAdDBtoFBAdPB(ad)

		adsPB = append(adsPB, newPB)
	}
	return adsPB
}

func convertFBAdDBtoFBAdPB(ad db.FacebookAd) *pb.FacebookAd {

	newPB := &pb.FacebookAd{
		AdId:                      fmt.Sprint(ad.AdID),
		PageId:                    fmt.Sprint(ad.PageID.Int64),
		PageName:                  ad.PageName.String,
		AdSnapshotUrl:             ad.AdSnapshotUrl.String,
		AdCreativeBody:            ad.AdCreativeBody.String,
		AdCreativeLinkCaption:     ad.AdCreativeLinkCaption.String,
		AdCreativeLinkDescription: ad.AdCreativeLinkDescription.String,
		AdCreativeLinkTitle:       ad.AdCreativeLinkTitle.String,
		AdDeliveryStartTime:       ad.AdDeliveryStartTime.String,
		AdDeliveryStopTime:        ad.AdDeliveryStopTime.String,
		FundingEntity:             ad.FundingEntity.String,
		ImpressionsMin:            ad.ImpressionsMin.Int32,
		ImpressionsMax:            ad.ImpressionsMax.Int32,
		SpendMin:                  ad.SpendMin.Int32,
		SpendMax:                  ad.SpendMax.Int32,
		Currency:                  ad.Currency.String,
		AdUrl:                     ad.AdUrl.String,
		SocialMediaFacebook:       ad.SocialMediaFacebook.String,
		SocialMediaInstagram:      ad.SocialMediaInstagram.String,
		SocialMediaWhatsapp:       ad.SocialMediaWhatsapp.String,
		SearchTerms:               ad.SearchTerms.String,
		AdCreationTime:            ad.AdCreationTime.String,
		PotentialReachMax:         ad.PotentialReachMax.Int32,
		PotentialReachMin:         ad.PotentialReachMin.Int32,
	}
	return newPB
}

type facebookJobListResponse struct {
	FacebookJob   *pb.CreateFacebookAdRequest `json:"facebook_job"`
	TotalFoundAds int64                       `json:"total_found_ads"`
	JobID         int64                       `json:"job_id"`
}

func convertJobDBtoJobResponse(jobs []db.FacebookJob) []*facebookJobListResponse {

	var jobsResponse []*facebookJobListResponse

	for _, job := range jobs {

		nJob := &pb.CreateFacebookAdRequest{
			AccessToken:        job.AccessToken.String,
			PageTotal:          int32(job.PageTotal.Int64),
			SearchTotal:        int32(job.SearchTotal.Int64),
			AdActiveStatus:     job.AdActiveStatus.String,
			AdDeliveryDateMax:  job.AdDeliveryDateMax.String,
			AdDeliveryDateMin:  job.AdDeliveryDateMin.String,
			AdReachedCountries: job.AdReachedCountries.String,
			SearchTerms:        job.SearchTerms.String,
		}
		response := facebookJobListResponse{
			FacebookJob:   nJob,
			JobID:         job.ID,
			TotalFoundAds: job.TotalFoundAds.Int64,
		}
		jobsResponse = append(jobsResponse, &response)
	}

	return jobsResponse
}

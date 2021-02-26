package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
	pb_etl "github.com/stanlee321/facebook-ads-server/pkg/etl/api/v1"
)

type getIndicatorABRequest struct {
	JobID int64 `uri:"job_id" binding:"required"`
}

func (server *Server) getIndicatorAB(ctx *gin.Context) {
	var req getIndicatorABRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ads, err := server.store.ListFacebookAdsByJobID(ctx,
		sql.NullInt64{Int64: req.JobID, Valid: true})

	log.Print(len(ads), req.JobID)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Prepare stream
	stream, err := server.facebookETLClient.CreateIndOne(ctx)

	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	// Send Stream
	for _, ad := range convertDBtoPBList(ads) {

		req := &pb_etl.CreateIndOneRequest{
			JobId:      req.JobID,
			FacebookAd: ad,
		}

		err = stream.Send(req)

		if err != nil {
			//log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			stream.RecvMsg(nil)
			return
		}
	}

	resIndAB, err := stream.CloseAndRecv()

	if err != nil {
		//log.Fatal("cannot receive response: ", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resIndAB)
}

func convertDBtoPBList(ads []db.FacebookAd) []*pb_etl.FacebookAd {

	var fbAdsPB []*pb_etl.FacebookAd

	for _, ad := range ads {
		fbAdsPB = append(fbAdsPB, convertAdDBtoAdPB(ad))
	}

	return fbAdsPB
}

func convertAdDBtoAdPB(ad db.FacebookAd) *pb_etl.FacebookAd {
	return &pb_etl.FacebookAd{
		AdId:                      fmt.Sprint(ad.AdID),
		PageId:                    fmt.Sprint(ad.PageID),
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
}

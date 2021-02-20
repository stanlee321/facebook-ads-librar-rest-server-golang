package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
	pb "github.com/stanlee321/facebook-ads-server/pkg/api/v1"
)

// Server serves HTTP for our facebook ads server
type Server struct {
	store          db.Store
	router         *gin.Engine
	facebookClient pb.FacebookAdsServiceClient
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store, facebookClient pb.FacebookAdsServiceClient) *Server {

	server := &Server{store: store, facebookClient: facebookClient}

	router := gin.Default()

	//router.GET("/api/facebook/ads/:id", server.getFacebookAd)
	router.GET("/api/facebook/ads/list/all/", server.listFacebookAds)
	router.GET("/api/facebook/ads/list/by_page_id", server.listFacebookAdsByPageID)
	router.GET("/api/facebook/ads/list/by_page_name", server.listFacebookAdsByPageName)
	router.POST("/api/facebook/ads/delete/:id", server.deleteFacebookAd)
	router.POST("/api/facebook/ads/create_job", server.createJob)
	router.GET("/api/facebook/jobs/list/all", server.listFacebookJobs)

	// Add routtes

	server.router = router

	return server
}

//Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

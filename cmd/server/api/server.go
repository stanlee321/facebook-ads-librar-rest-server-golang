package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
	pb "github.com/stanlee321/facebook-ads-server/pkg/ads/api/v1"
	pb_etl "github.com/stanlee321/facebook-ads-server/pkg/etl/api/v1"
)

// Server serves HTTP for our facebook ads server
type Server struct {
	store             db.Store
	router            *gin.Engine
	facebookClient    pb.FacebookAdsServiceClient
	facebookETLClient pb_etl.FacebookAdsETLServiceClient
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store,
	facebookClient pb.FacebookAdsServiceClient,
	facebookETLClient pb_etl.FacebookAdsETLServiceClient) *Server {

	server := &Server{store: store,
		facebookClient:    facebookClient,
		facebookETLClient: facebookETLClient,
	}

	router := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	//router.GET("/api/facebook/ads/:id", server.getFacebookAd)
	router.GET("/api/facebook/ads/list/all/", server.listFacebookAds)
	router.GET("/api/facebook/ads/list/by_page_id", server.listFacebookAdsByPageID)
	router.GET("/api/facebook/ads/list/by_page_name", server.listFacebookAdsByPageName)
	router.POST("/api/facebook/ads/delete/:id", server.deleteFacebookAd)
	router.POST("/api/facebook/ads/create_job", server.createJob)
	router.GET("/api/facebook/jobs/list/all", server.listFacebookJobs)

	// ETLs
	router.GET("/api/facebook/ads/etl/ind_a_b/:job_id", server.getIndicatorAB)

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

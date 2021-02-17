package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
)

// Server serves HTTP for our facebook ads server
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {

	server := &Server{store: store }

	router := gin.Default()

	router.POST("/api/facebook/ads/create", server.createFacebookAd)
	router.GET("/api/facebook/ads/:id", server.getFacebookAd)
	router.GET("/api/facebook/ads", server.listFacebookAds)
	router.GET("/api/facebook/ads/by_page_id", server.listFacebookAdsByPageID)
	router.GET("/api/facebook/ads/by_page_name", server.listFacebookAdsByPageName)
	router.GET("/api/facebook/ads/by_page_ad_id", server.listFacebookAdsByAdID)
	router.POST("/api/facebook/ads/delete/:id", server.deleteFacebookAd)

	// Add routtes

	server.router = router
	
	return server
}


//Start runs the HTTP server on a specific address
func (server *Server ) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H{
	return gin.H{
		"error": err.Error(),
	}
}
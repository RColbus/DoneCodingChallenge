package api

import (
	db "doneAPI/db/sqlc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Server serves HTTP requests for our API server
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "PUT", "PATCH", "POST"}
	corsConfig.AllowHeaders = []string{"Origin"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	//OPTIONS request for ReactJS
	corsConfig.AddAllowHeaders("OPTIONS")

	// Register the middleware
	router.Use(cors.New(corsConfig))

	router.GET("api/v1", server.healthCheck)
	router.POST("api/v1/registrations", server.createRegistration)
	router.GET("api/v1/registrations/:id", server.getRegistration)
	router.GET("api/v1/registrations", server.listRegistration)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse returns http error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// HealthCheckResponse healthCheck respone object
type HealthCheckResponse struct {
	Note string `json:"note" binding:"required"`
}

// healthCheck Simple server health check response to verify server is running
func (server *Server) healthCheck(ctx *gin.Context) {
	resp := HealthCheckResponse{
		Note: "Welcome to the Done API service Homepage",
	}

	ctx.JSON(http.StatusOK, resp)
}

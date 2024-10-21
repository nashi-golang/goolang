package routes

import (
	"github.com/gin-gonic/gin"
	"goolang/handlers"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()

	// wedding routes
	weddingGroup := router.Group("/weddings")
	handlers.RouteWeddingHandler(weddingGroup)

	// guest routes
	guestGroup := router.Group("/guests")
	handlers.RouteGuestHandler(guestGroup)

	return router
}

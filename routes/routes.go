package routes

import (
	"github.com/gin-gonic/gin"
	"goolang/handlers"
)

func SetRoutes(address string) {
	router := gin.Default()

	router.GET("/weddings", handlers.GetWeddings)
	router.GET("/weddings/:guid", handlers.GetWedding)
	router.POST("/weddings", handlers.CreateWeddingRequest)
	router.PUT("/weddings/:guid", handlers.UpdateWedding)
	router.DELETE("/weddings/:guid", handlers.DeleteWedding)

	router.Run(address)
}

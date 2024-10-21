package handlers

import (
	"github.com/gin-gonic/gin"
	"goolang/models"
	"goolang/util"
	"log"
	"time"
)

type WeddingDto struct {
	Guid          string    `json:"guid"`
	StartDatetime string    `json:"start_datetime" binding:"required"`
	Location      string    `json:"location" binding:"required"`
	Groom         PersonDto `json:"groom"`
	Bride         PersonDto `json:"bride"`
}

func RouteWeddingHandler(router *gin.RouterGroup) {
	router.POST("", CreateWeddingRequest)
	router.GET("", GetWeddings)
	router.GET("/:guid", GetWedding)
	router.PUT("/:guid", UpdateWedding)
	router.DELETE("/:guid", DeleteWedding)
}

func CreateWeddingRequest(c *gin.Context) {
	var newWedding WeddingDto
	if err := c.ShouldBindJSON(&newWedding); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	guid := util.GenerateGuid()
	tartDateTime, err := time.Parse("20060102150405", newWedding.StartDatetime)
	if err != nil {
		log.Fatal("failed to parse time:", err)
	}

	newGroomItem := models.Person{
		Guid:        util.GenerateGuid(),
		WeddingGuid: guid,
		Name:        newWedding.Groom.Name,
		PhoneNumber: newWedding.Groom.PhoneNumber}

	newBrideItem := models.Person{
		Guid:        util.GenerateGuid(),
		WeddingGuid: guid,
		Name:        newWedding.Bride.Name,
		PhoneNumber: newWedding.Bride.PhoneNumber}

	newWeddingItem := models.Wedding{
		Guid:          guid,
		StartDatetime: tartDateTime,
		Location:      newWedding.Location,
		Groom:         newGroomItem.Guid,
		Bride:         newBrideItem.Guid}

	models.DB.Create(newGroomItem)
	models.DB.Create(newBrideItem)
	models.DB.Create(newWeddingItem)

	c.JSON(201, newWeddingItem)
}

func GetWeddings(c *gin.Context) {
	var weddings []models.Wedding
	models.DB.Find(&weddings)

	c.JSON(200, weddings)
}

func GetWedding(c *gin.Context) {
	guid := c.Param("guid")
	var wedding models.Wedding
	models.DB.First(&wedding, "guid = ?", guid)

	if wedding.Guid == "" {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(200, wedding)
}

func UpdateWedding(c *gin.Context) {
	guid := c.Param("guid")
	var wedding models.Wedding
	models.DB.First(&wedding, "guid = ?", guid)

	if wedding.Guid == "" {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	var updatedWedding WeddingDto
	if err := c.ShouldBindJSON(&updatedWedding); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tartDateTime, err := time.Parse("20060102150405", updatedWedding.StartDatetime)
	if err != nil {
		log.Fatal("failed to parse time:", err)
	}

	models.DB.Model(&wedding).Updates(
		models.Wedding{
			StartDatetime: tartDateTime,
			Location:      updatedWedding.Location,
			Groom:         updatedWedding.Groom.Guid,
			Bride:         updatedWedding.Bride.Guid})

	c.JSON(200, updatedWedding)
}

func DeleteWedding(c *gin.Context) {
	guid := c.Param("guid")
	var wedding models.Wedding
	models.DB.First(&wedding, "guid = ?", guid)

	if wedding.Guid == "" {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&wedding)

	c.JSON(204, gin.H{})
}

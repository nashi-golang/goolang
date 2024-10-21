package handlers

import (
	"github.com/gin-gonic/gin"
	"goolang/models"
	"goolang/util"
)

type PersonDto struct {
	Guid        string `json:"guid"`
	WeddingGuid string `json:"wedding_guid"`
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func RouteGuestHandler(group *gin.RouterGroup) {
	group.POST("", createGuestQuest)
	group.GET("", getGuests)
	group.GET("/:guid", getGuest)
	group.PUT("/:guid", updateGuest)
	group.DELETE("/:guid", deleteGuestRequest)
	group.GET("/wedding/:weddingGuid", getGuestByWedding)
}

func createGuestQuest(c *gin.Context) {
	var newPerson PersonDto
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	guid := util.GenerateGuid()

	newPersonItem := models.Person{
		Guid:        guid,
		WeddingGuid: newPerson.WeddingGuid,
		Name:        newPerson.Name,
		PhoneNumber: newPerson.PhoneNumber,
	}

	models.DB.Create(newPersonItem)
	c.JSON(201, newPersonItem)
}

func deleteGuestRequest(c *gin.Context) {
	guid := c.Param("guid")
	models.DB.Delete(&models.Person{}, guid)
	c.Status(204)
}

func getGuests(c *gin.Context) {
	var guests []models.Person
	models.DB.Find(&guests)
	c.JSON(200, guests)
}

func getGuest(c *gin.Context) {
	guid := c.Param("guid")
	var guest models.Person
	models.DB.First(&guest, guid)
	c.JSON(200, guest)
}

func updateGuest(c *gin.Context) {
	var updatedPerson PersonDto
	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	guid := c.Param("guid")

	var person models.Person
	models.DB.First(&person, guid)

	person.Name = updatedPerson.Name
	person.PhoneNumber = updatedPerson.PhoneNumber

	models.DB.Save(&person)
	c.JSON(200, person)
}

func getGuestByWedding(c *gin.Context) {
	weddingGuid := c.Param("weddingGuid")
	var guests []models.Person
	models.DB.Find(&guests, "wedding_guid = ?", weddingGuid)
	c.JSON(200, guests)
}

package controllers

import (
	"net/http"
	"mix3d/internal/models"

	"github.com/gin-gonic/gin"
)

// type PersonController struct {
// 	PersonRepository data.PersonRepository
// }

// func NewPersonController(PersonRepository firestore.PersonRepository) *PersonController {
// 	return &PersonController{PersonRepository: PersonRepository}
// }

func GetPeople(c *gin.Context) {
	People := models.People(c)

	c.JSON(http.StatusOK, People)

}

func GetPerson(c *gin.Context) {
	Person := models.PersonById(c.Param("id"))

	if Person.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Person found"})
		return
	}
	c.JSON(http.StatusOK, Person)
}

func CreatePerson(c *gin.Context) {
	var Person models.Person

	if err := c.ShouldBindJSON(&Person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err})
		return
	}
	result := models.CreatePerson(Person)

	if result["status"] == "error" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, result)
}

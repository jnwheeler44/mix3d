package controllers

import (
	"net/http"

	"parksport-go/internal/models"

	"github.com/gin-gonic/gin"
)

func GetEntities(c *gin.Context) {
	entities := models.Entities(c)

	c.JSON(http.StatusOK, entities)

}

func GetEntity(c *gin.Context) {
	entity := models.EntityById(c.Param("id"))

	if entity.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No entity found!"})
		return
	}
	c.JSON(http.StatusOK, entity)
}

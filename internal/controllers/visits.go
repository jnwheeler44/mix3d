package controllers

import (
	"net/http"

	"parksport-go/internal/models"

	"github.com/gin-gonic/gin"
)

func GetVisits(c *gin.Context) {
	users := models.Visits(c)

	c.JSON(http.StatusOK, users)

}

func GetVisit(c *gin.Context) {
	user := models.VisitById(c.Param("id"))

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No visit found!"})
		return
	}
	c.JSON(http.StatusOK, user)
}

package controllers

import (
	"ginBestPractice/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MakananController struct {
	service services.MakananService
}

func NewMakananController(service services.MakananService) *MakananController {
	return &MakananController{service}
}

func (mkn *MakananController) GetById(c *gin.Context) {
	idString := c.Param("id")
	idUint, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	makanan, err := mkn.service.GetById(uint(idUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "Makanan not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, makanan)

}

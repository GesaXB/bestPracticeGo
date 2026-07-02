package controllers

import (
	"ginBestPractice/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IpkController struct {
	service services.IpkService
}

func NewIpkController(service services.IpkService) *IpkController {
	return &IpkController{service}
}

func (i IpkController) GetAllIpk(c *gin.Context) {
	ipk, err := i.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ipk)
}

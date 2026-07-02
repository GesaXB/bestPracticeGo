package controllers

import (
	"ginBestPractice/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MahasiswaController struct {
	service services.MahasiswaService
}

func NewMahasiswaController(service services.MahasiswaService) *MahasiswaController {
	return &MahasiswaController{service}
}

func (m *MahasiswaController) GetAllMahasiswa(c *gin.Context) {
	mhs, err := m.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mhs)
}

func (m *MahasiswaController) FindById(c *gin.Context) {
	nimStr := c.Param("nim")

	nimUint, err := strconv.ParseUint(nimStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format NIM tidak valid, harus berupa angka"})
		return
	}

	mhs, err := m.service.GetByID(uint(nimUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "data tidak bisa ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mhs)
}

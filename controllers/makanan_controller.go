package controllers

import (
	"ginBestPractice/dtos"
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

func (mkn *MakananController) GetAll(c *gin.Context) {
	listMakanan, err := mkn.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, listMakanan)
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

func (mkn *MakananController) Create(c *gin.Context) {
	var input dtos.MakananCreateRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := mkn.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan ke database"})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (mkn MakananController) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus berupa angka"})
		return
	}

	var input dtos.MakananUpdateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := mkn.service.Update(uint(id), input)
	if err != nil {
		if err.Error() == "makanan tidak bisa ditemukan" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

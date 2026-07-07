package controllers

import (
	"ginBestPractice/dtos"
	"ginBestPractice/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KategoriController struct {
	service services.KategoriService
}

func NewKategoriController(service services.KategoriService) *KategoriController {
	return &KategoriController{service}
}

func (k *KategoriController) GetAll(c *gin.Context) {
	kategoris, err := k.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kategoris)
}

func (k *KategoriController) GetById(c *gin.Context) {
	idString := c.Param("id")
	idUint, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	res, err := k.service.GetById(uint(idUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(400, map[string]string{
				"error": "kategori tidak ditemukan",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (k KategoriController) Create(c *gin.Context) {
	var input dtos.KategoriCreateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := k.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan ke database"})
		return
	}

	c.JSON(http.StatusCreated, res)
}

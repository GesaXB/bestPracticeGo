package controllers

import (
	"ginBestPractice/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{service}
}

func (bc BookController) GetAllBooks(c *gin.Context) {
	book, err := bc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc BookController) GetById(c *gin.Context) {
	idString := c.Param("id")
	idUint, err := strconv.ParseUint(idString, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "format id tidak valid, harus angka!"})
		return
	}

	book, err := bc.service.GetByID(uint(idUint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)

}

package routes

import (
	"ginBestPractice/controllers"
	"ginBestPractice/repositories"
	"ginBestPractice/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Mahasiswa
	mhsRepo := repositories.NewMahasiswaRepository(db)
	mhsService := services.NewMahasiswaService(mhsRepo)
	mhsController := controllers.NewMahasiswaController(mhsService)

	// Ipk
	ipkRepo := repositories.NewIpkRepository(db)
	ipkService := services.IpkService(ipkRepo)
	ipkController := controllers.NewIpkController(ipkService)

	// Book
	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)

	// Makanan
	makananRepo := repositories.NewMakananRepository(db)
	makananService := services.NewMakananService(makananRepo)
	makananController := controllers.NewMakananController(makananService)

	api := r.Group("/api")
	{
		api.GET("/mahasiswa", mhsController.GetAllMahasiswa)
		api.GET("/mahasiswa/:nim", mhsController.FindById)
		api.GET("/ipk", ipkController.GetAllIpk)
		api.GET("/books", bookController.GetAllBooks)
		api.GET("/book/:id", bookController.GetById)
		api.GET("/makanan/:id", makananController.GetById)
	}
}

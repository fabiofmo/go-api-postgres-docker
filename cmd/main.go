package main

import (
	"net/http"

	"github.com/fabiofmo/go-api-postgres-docker/controller"
	"github.com/fabiofmo/go-api-postgres-docker/db"
	"github.com/fabiofmo/go-api-postgres-docker/repository"
	"github.com/fabiofmo/go-api-postgres-docker/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)   // camada de respository
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)       // camada de usecase
	ProductController := controller.NewProductController(ProductUsecase) // camada de controllers

	router.GET("/product", ProductController.GetProducts)
	router.GET("/product/:productId", ProductController.GetProductById)
	router.POST("/product/add", ProductController.CreateProduct)

	router.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })
	router.Run("localhost:8080")

}

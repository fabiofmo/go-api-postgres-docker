package main

import (
	"net/http"

	"github.com/fabiofmo/go-api-postgres-docker/controller"
	"github.com/fabiofmo/go-api-postgres-docker/model"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	ProductController := controller.NewProductController()
	router.GET("/products", ProductController.GetProducts)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/product", getAllProducts)
	router.Run("localhost:8080")

}

func getAllProducts(context *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}
	context.IndentedJSON(http.StatusOK, products)
}

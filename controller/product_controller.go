package controller

import (
	"net/http"

	"github.com/fabiofmo/go-api-postgres-docker/model"
	"github.com/fabiofmo/go-api-postgres-docker/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(context *gin.Context) {

	products, err := p.productUsecase.GetProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(context *gin.Context) {

	var product model.Product
	err := context.BindJSON(&product)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, insertedProduct)
}

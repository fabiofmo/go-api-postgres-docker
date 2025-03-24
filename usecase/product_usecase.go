package usecase

/* service
 */

import (
	"github.com/fabiofmo/go-api-postgres-docker/model"
	"github.com/fabiofmo/go-api-postgres-docker/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (int, error) {

	productId, err := pu.CreateProduct(product)

	if err != nil {
		return 0, err
	}

	return productId, nil
}

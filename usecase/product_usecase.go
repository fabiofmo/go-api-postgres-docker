package usecase

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

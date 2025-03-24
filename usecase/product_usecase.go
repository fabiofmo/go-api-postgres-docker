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

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	newId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = newId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(idProduct int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(idProduct)

	if err != nil {
		return nil, err
	}

	return product, nil

}

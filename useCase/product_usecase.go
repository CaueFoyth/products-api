package usecase

import (
	"products-api/model"
	"products-api/repository"
)

type ProductsUseCase struct {
	repository repository.ProductRepository
}

func NewProductsUseCase(repo repository.ProductRepository) ProductsUseCase {
	return ProductsUseCase{
		repository: repo,
	}
}

func (pu *ProductsUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductsUseCase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = id
	return product, nil
}

func (pu *ProductsUseCase) GetProductByID(id int) (model.Product, error) {
	return pu.repository.GetProductByID(id)
}

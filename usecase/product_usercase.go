package usecase

import (
	"go_mod_api/model"
	"go_mod_api/repository"
)

type ProductUserCase struct {
	repository repository.ProductRepository
}

func NewProductUserCase(repo repository.ProductRepository) ProductUserCase {
	return ProductUserCase{
		repository: repo,
	}
}

func (pu *ProductUserCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUserCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.Id = productId

	return product, nil
}

func (pu *ProductUserCase) GetProductByID(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

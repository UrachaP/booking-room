package productservice

import (
	"errors"
	"log"

	"bookingrooms/internal/models"
	productrepository "bookingrooms/internal/repositories/product"
)

type service struct {
	repository productrepository.ProductRepository
}

type ProductService interface {
	AddProductStock(product models.Product) error
	CutProductStock(requestProduct models.RequestProduct) error
}

func NewProductService(repository productrepository.ProductRepository) ProductService {
	return service{repository: repository}
}

func (s service) AddProductStock(product models.Product) error {
	return s.repository.CreateProduct(product)
}

func (s service) CutProductStock(requestProduct models.RequestProduct) error {
	product, err := s.repository.GetProduct(requestProduct.ID)
	if err != nil {
		return err
	}

	product.Amount = product.Amount - requestProduct.Amount
	if product.Amount < 0 {
		return errors.New("out of stock")
	}
	log.Println(product)

	return s.repository.UpdateProduct(product)
}

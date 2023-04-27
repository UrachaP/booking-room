package productrepository

import (
	"bookingrooms/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type ProductRepository interface {
	CreateProduct(product models.Product) error
	GetProduct(id int) (models.Product, error)
	UpdateProduct(product models.Product) error
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return repository{db: db}
}

func (r repository) CreateProduct(product models.Product) error {
	return r.db.Create(&product).Error
}

func (r repository) GetProduct(id int) (models.Product, error) {
	var product models.Product
	return product, r.db.Find(&product, id).Error
}

func (r repository) UpdateProduct(product models.Product) error {
	return r.db.Save(&product).Error
}

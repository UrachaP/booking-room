package productrepository

import (
	"testing"

	"bookingrooms/internal/models"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRepository_CreateProduct(t *testing.T) {
	product := models.Product{Name: "apple", Amount: 3}

	db, _ := gorm.Open(mysql.Open("test:12345678@tcp(203.154.71.142:3306)/exam?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	r := NewProductRepository(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err := r.CreateProduct(product)

	assert.Equal(t, nil, err)
}

func TestRepository_GetProduct(t *testing.T) {
	expected := models.Product{ID: 1, Name: "apple", Amount: 3}
	id := 1

	db, _ := gorm.Open(mysql.Open("test:12345678@tcp(203.154.71.142:3306)/exam?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	r := NewProductRepository(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	actual, err := r.GetProduct(id)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestRepository_UpdateProduct(t *testing.T) {
	product := models.Product{ID: 2, Name: "apple", Amount: 0}

	db, _ := gorm.Open(mysql.Open("test:12345678@tcp(203.154.71.142:3306)/exam?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	r := NewProductRepository(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err := r.UpdateProduct(product)

	assert.Equal(t, nil, err)
}

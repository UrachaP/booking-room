package roomrepository

import (
	"bookingrooms/internal/models"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_DeletedRoom(t *testing.T) {
	id := 11
	room := models.Rooms{
		ID: id,
	}
	expected := int64(0)

	DB, _ := gorm.Open(mysql.Open("test:12345678@tcp(203.154.71.142:3306)/exam?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	r := Repository{
		DB: DB,
	}
	r.DB.Create(&models.Rooms{ID: id})
	err := r.DeleteRoom(&room)
	actual := r.DB.First(&models.Rooms{}, id).RowsAffected

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

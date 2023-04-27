package userrepository

import (
	"errors"
	"testing"
	"time"

	"bookingrooms/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRepository_Register(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewUserRepository(db, nil)

	t.Run("success", func(t *testing.T) {
		expected := models.Users{
			Model:        models.Model{ID: 100},
			Username:     "username_test_register",
			PasswordHash: "$2a$14$BU5cZk1dgGuXb0N0yAyvkehIqzQQmqyPWcVCcDArkiegj3WIdxVdy",
			CreatedBy:    100,
		}
		//prepare data
		user := models.Users{Model: models.Model{ID: 100}, Username: "username_test_register", PasswordHash: "$2a$14$BU5cZk1dgGuXb0N0yAyvkehIqzQQmqyPWcVCcDArkiegj3WIdxVdy"}
		var actualDB models.Users
		//actual
		err := r.Register(user)
		db.Where("username = ?", user.Username).First(&actualDB)
		actual := models.Users{Model: models.Model{ID: actualDB.ID}, Username: actualDB.Username, PasswordHash: actualDB.PasswordHash, CreatedBy: actualDB.CreatedBy}
		//result
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
		assert.WithinDuration(t, time.Now(), actualDB.CreatedAt, time.Second)
		assert.WithinDuration(t, time.Now(), actualDB.UpdatedAt, time.Second)
		//clear data
		db.Unscoped().Where("username = ?", user.Username).Delete(user)
	})

	t.Run("error username is duplicated", func(t *testing.T) {
		//prepare data
		user := models.Users{Username: "moji1", PasswordHash: ""}
		//expected
		expected := errors.New("username is duplicated")
		//actual
		actual := r.Register(user)
		//result
		assert.Equal(t, expected, actual)
	})

	t.Run("error created and rollback", func(t *testing.T) {
		actual := r.Register(models.Users{Username: "username_test", AGrade: "ERROR"})
		assert.NotEqual(t, nil, actual)

		count := db.Where("username = 'username_test'").First(&models.Users{}).RowsAffected
		assert.Equal(t, int64(0), count)
	})
}

func Test_GetPasswordHash(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewUserRepository(db, nil)

	t.Run("success", func(t *testing.T) {
		//assert
		expected := models.Users{
			Model:        models.Model{ID: 1},
			PasswordHash: "$2a$14$BU5cZk1dgGuXb0N0yAyvkehIqzQQmqyPWcVCcDArkiegj3WIdxVdy",
		}
		//actual
		actual, err := r.GetPasswordHash("moji1")
		//result
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("error record not found", func(t *testing.T) {
		//assert
		expected := models.Users{}
		expectedErr := "record not found"
		//actual
		actual, err := r.GetPasswordHash("username not found")
		//result
		assert.Equal(t, expectedErr, err.Error())
		assert.Equal(t, expected, actual)
	})

}

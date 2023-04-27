package tempimagerepository

import (
	"errors"
	"testing"
	"time"

	"bookingrooms/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_ValidateTempImages_Not_Have_Temp_Image(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		id := 0
		expected := errors.New("not have temp image")
		actual := r.HaveTempImages(id)
		assert.Equal(t, expected, actual)
	})

}

func Test_ValidateTempImages_DeletedAt(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		expected := errors.New("not have temp image")
		id := 1

		actual := r.HaveTempImages(id)

		assert.Equal(t, expected, actual)
	})

	t.Run("have not deleted yet", func(t *testing.T) {
		id := 2

		actual := r.HaveTempImages(id)

		assert.Equal(t, nil, actual)
	})

}

func Test_UpdateTempImages(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		expected := "0"
		id := 3

		_ = r.UpdateTempImages(&models.TempImages{Model: models.Model{ID: id}, Temp: "0"})
		var actual string
		db.Model(&models.TempImages{}).Select("temp").Where("id = ?", id).First(&actual)

		assert.Equal(t, expected, actual)
	})
}

func Test_SaveTempImages(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		expected := models.TempImages{
			Model:     models.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Temp:      "1",
			Path:      "created",
			Name:      "created",
			Extension: "test",
		}
		tempImages := models.TempImages{
			Model:     models.Model{ID: 4},
			Path:      "created",
			Name:      "created",
			Extension: "test",
		}

		var actual models.TempImages
		err := r.SaveTempImages(tempImages)
		db.First(&actual, tempImages.ID)

		assert.Equal(t, nil, err)
		assert.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
		assert.WithinDuration(t, expected.UpdatedAt, actual.UpdatedAt, time.Second)
		actual.CreatedAt = expected.CreatedAt
		actual.UpdatedAt = expected.UpdatedAt
		assert.Equal(t, expected, actual)

		db.Unscoped().Delete(&tempImages)
	})

}

func Test_SoftDeletedAtTempImages(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		expected := time.Now()
		tempImage := models.TempImages{Model: models.Model{ID: 5}}

		_ = r.SoftDeletedAtTempImages(&tempImage)
		var actual models.TempImages
		db.Unscoped().Select("deleted_at").Where("id = ?", tempImage.ID).First(&actual)

		assert.WithinDuration(t, expected, actual.DeletedAt.Time, time.Second)
	})

}

func Test_GetTempImagesById(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		id := 6
		expected := models.TempImages{
			Model:     models.Model{ID: id},
			Path:      "get",
			Name:      "get",
			Extension: "test",
			Temp:      "1",
		}

		actual, _ := r.GetTempImagesById(id)

		assert.Equal(t, expected, actual)
	})

	t.Run("deleted", func(t *testing.T) {
		expected := models.TempImages{}
		id := 1

		actual, err := r.GetTempImagesById(id)

		assert.Equal(t, expected, actual)
		assert.Error(t, err, "record not found")
	})
}

func Test_GetTempImagesIsTempMoreOneDay(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		expected := []models.TempImages{
			{
				Model: models.Model{
					ID:        7,
					CreatedAt: time.Date(2022, 11, 11, 0, 0, 0, 0, time.Local),
				},
				Path:      "get",
				Name:      "get",
				Extension: "test",
				Temp:      "1",
			},
		}

		actual, _ := r.GetTempImagesIsTmpMoreOneDay()

		assert.Equal(t, expected, actual)
	})
}

func Test_GetTempImagesIsDeletedMoreOneDay(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		expected := []models.TempImages{
			{
				Path:      "deleted",
				Name:      "deleted",
				Extension: "test",
				Temp:      "1",
				Model: models.Model{
					ID: 1,
					DeletedAt: gorm.DeletedAt{
						Time:  time.Date(2023, 01, 1, 0, 0, 0, 0, time.Local),
						Valid: true,
					},
				},
			},
		}

		actual, _ := r.GetTempImagesIsDeletedMoreOneDay()

		assert.Equal(t, expected, actual)
	})

}

func Test_GetTempImagesIsNotTempNotUsedMoreOneDay(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		expected := []models.TempImages{
			{
				Path:      "updated",
				Name:      "updated",
				Extension: "test",
				Temp:      "0",
				Model: models.Model{
					ID:        8,
					UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
		}

		actual, _ := r.GetTempImagesIsNotTmpNotUsedMoreOneDay()

		assert.Equal(t, expected, actual)
	})

}

func Test_DeleteFileImages(t *testing.T) {
	//prepare data
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	r := NewTempImageRepository(db)

	t.Run("success", func(t *testing.T) {
		id := 9
		tempImage := []models.TempImages{
			models.TempImages{Model: models.Model{ID: id}},
		}
		expected := 0

		r.DeleteFileImages(tempImage)
		var actual int
		db.Raw("SELECT COUNT(*) FROM temp_images WHERE id = ?", id).First(&actual)

		assert.Equal(t, expected, actual)
	})
}

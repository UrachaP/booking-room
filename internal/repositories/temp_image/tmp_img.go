package tempimagerepository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"bookingrooms/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewTempImageRepository(db *gorm.DB) TempImageRepository {
	return repository{db: db}
}

type TempImageRepository interface {
	HaveTempImages(id int) error
	UpdateTempImages(tempImage *models.TempImages) error
	GetTempImagesIsTmpMoreOneDay() ([]models.TempImages, error)
	GetTempImagesIsDeletedMoreOneDay() ([]models.TempImages, error)
	GetTempImagesIsNotTmpNotUsedMoreOneDay() ([]models.TempImages, error)
	DeleteFileImages(tempImages []models.TempImages)
	SaveTempImages(tempImages models.TempImages) error
	SoftDeletedAtTempImages(tempImages *models.TempImages) error
	GetTempImagesById(id int) (models.TempImages, error)
}

// HaveTempImages duplicate(have image) and not deleted
// when query deleted use Unscoped()
func (r repository) HaveTempImages(id int) error {
	var count int
	err := r.db.
		Model(&models.TempImages{}).
		Select("COUNT(*)").
		Where("id = ?", id).First(&count).
		Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("not have temp image")
	}
	return nil
}

func (r repository) UpdateTempImages(tempImage *models.TempImages) error {
	return r.db.Updates(tempImage).Error
}

func (r repository) SaveTempImages(tempImages models.TempImages) error {
	return r.db.Create(&tempImages).Error
}

func (r repository) SoftDeletedAtTempImages(tempImages *models.TempImages) error {
	return r.db.Delete(tempImages).Error
}

func (r repository) GetTempImagesById(id int) (models.TempImages, error) {
	var tmpImg models.TempImages
	return tmpImg, r.db.
		Where("id = ?", id).
		First(&tmpImg).
		Error
}

func (r repository) GetTempImagesIsTmpMoreOneDay() ([]models.TempImages, error) {
	var TempImages []models.TempImages
	err := r.db.
		Where("temp = '1'").
		Where("DATE_SUB(NOW(),INTERVAL 1 DAY)>created_at").
		Find(&TempImages).Error
	return TempImages, err
}

func (r repository) GetTempImagesIsDeletedMoreOneDay() ([]models.TempImages, error) {
	var TempImages []models.TempImages
	return TempImages, r.db.Unscoped().
		Where("DATE_SUB(NOW(),INTERVAL 1 DAY)>deleted_at").
		//Where("deleted_at != 0").
		Find(&TempImages).
		Error
}

func (r repository) GetTempImagesIsNotTmpNotUsedMoreOneDay() ([]models.TempImages, error) {
	var TempImages []models.TempImages
	err := r.db.Debug().
		Select("temp_images.*").
		Joins("LEFT JOIN rooms_images ON rooms_images.image_id = temp_images.id").
		Joins("LEFT JOIN users ON users.image_id = temp_images.id ").
		Where(r.db.
			Where("temp = '0'").
			Where("temp_images.updated_at < NOW() - INTERVAL 1 DAY")).
		Where(r.db.
			Where("rooms_images.image_id IS NULL").
			Where("users.image_id IS NULL")).
		Find(&TempImages).Error
	return TempImages, err
}

func (r repository) DeleteFileImages(tempImages []models.TempImages) {
	for _, image := range tempImages {
		err := os.Remove(image.Path)
		if err != nil {
			// log and continue shouldn't stop running
			log.Println(fmt.Sprintf("Error : %s at image ID: %d", err.Error(), image.ID))
		}
		// delete db
		err = r.DeletedTempImagesById(&image)
		if err != nil {
			log.Println(fmt.Sprintf("Error : %s at image ID: %d", err.Error(), image.ID))
		}
	}
}

func (r repository) DeletedTempImagesById(tempImage *models.TempImages) error {
	return r.db.Unscoped().Delete(tempImage).Error
}

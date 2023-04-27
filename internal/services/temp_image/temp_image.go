package tempimageservice

import (
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"bookingrooms/internal/models"
	tempimagerepository "bookingrooms/internal/repositories/temp_image"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type service struct {
	repository tempimagerepository.TempImageRepository
}

func NewTempImageService(repository tempimagerepository.TempImageRepository) service {
	return service{repository: repository}
}

type TempImageService interface {
	NewTempImage(id int) *models.TempImages
	ReadFileImage(file *multipart.FileHeader) (models.TempImages, error)
	ResizeFileImage(file *multipart.FileHeader, dst *os.File) error
	HaveTempImages(id int) error
	UpdateTempImages(tempImage *models.TempImages) error
	DeleteTempImagesIsTmpMoreOneDay() error
	DeleteTempImagesIsDeletedMoreOneDay() error
	DeleteTempImagesIsNotTmpNotUsedMoreOneDay() error
	GetTempImage(id int) (models.TempImages, error)
	SaveTempImage(images models.TempImages) error
	SoftDeletedAtTempImages(id int) error
	DisplayImage(id int) (string, error)
}

func (s service) NewTempImage(id int) *models.TempImages {
	return &models.TempImages{
		Model: models.Model{
			ID: id,
		},
		Temp: "0",
	}
}

func (s service) ReadFileImage(file *multipart.FileHeader) (models.TempImages, error) {
	// Destination
	filePath := filepath.Join(filepath.Join("assets/image/", uuid.NewString()))
	dst, err := os.Create(filePath)
	if err != nil {
		return models.TempImages{}, err
	}
	defer dst.Close()

	if file.Size >= 1000000 {
		s.ResizeFileImage(file, dst)
	}

	// get root path
	cwd, err := os.Getwd()
	if err != nil {
		return models.TempImages{}, err
	}

	tempImage := models.TempImages{
		Path:      cwd + "/" + filePath,
		Name:      file.Filename,
		Extension: strings.Replace(filepath.Ext(file.Filename), ".", "", -1),
	}
	return tempImage, err
}

func (s service) ResizeFileImage(file *multipart.FileHeader, dst *os.File) error {
	fileType := filepath.Ext(file.Filename)

	// open cause resize file
	src, err := file.Open()
	if err != nil {
		return err
	}

	// decode jpeg into image.Image
	// check file type
	image, err := s.decodeImage(src, fileType)

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	imageResize := resize.Resize(1000, 0, image, resize.Lanczos3)

	err = s.encodeImage(dst, fileType, imageResize)
	if err != nil {
		return err
	}

	defer src.Close()

	return nil
}

func (s service) decodeImage(src multipart.File, fileType string) (image.Image, error) {
	var image image.Image
	var err error
	if fileType == ".jpeg" {
		image, err = jpeg.Decode(src)
	} else if fileType == ".png" {
		image, err = png.Decode(src)
	}
	return image, err
}

func (s service) encodeImage(file *os.File, fileType string, imageResize image.Image) error {
	var err error
	if fileType == ".jpeg" {
		err = jpeg.Encode(file, imageResize, nil)
	} else if fileType == ".png" {
		err = png.Encode(file, imageResize)
	}
	return err
}

func (s service) HaveTempImages(id int) error {
	return s.repository.HaveTempImages(id)
}

func (s service) UpdateTempImages(tempImage *models.TempImages) error {
	return s.repository.UpdateTempImages(tempImage)
}

func (s service) DeleteTempImagesIsTmpMoreOneDay() error {
	images, err := s.repository.GetTempImagesIsDeletedMoreOneDay()
	if err != nil {
		return err
	}
	s.repository.DeleteFileImages(images)
	return nil
}

func (s service) DeleteTempImagesIsDeletedMoreOneDay() error {
	images, err := s.repository.GetTempImagesIsDeletedMoreOneDay()
	if err != nil {
		return err
	}
	s.repository.DeleteFileImages(images)
	return nil
}

func (s service) DeleteTempImagesIsNotTmpNotUsedMoreOneDay() error {
	images, err := s.repository.GetTempImagesIsNotTmpNotUsedMoreOneDay()
	if err != nil {
		return err
	}
	s.repository.DeleteFileImages(images)
	return nil
}

func (s service) GetTempImage(id int) (models.TempImages, error) {
	return s.repository.GetTempImagesById(id)
}

func (s service) SaveTempImage(images models.TempImages) error {
	return s.repository.SaveTempImages(images)
}

func (s service) SoftDeletedAtTempImages(id int) error {
	return s.repository.SoftDeletedAtTempImages(s.NewTempImage(id))
}

func (s service) DisplayImage(id int) (string, error) {
	err := s.repository.HaveTempImages(id)
	if err != nil {
		return "", err
	}
	tempImage, err := s.repository.GetTempImagesById(id)
	if err != nil {
		return "", err
	}
	return tempImage.Path, nil
}

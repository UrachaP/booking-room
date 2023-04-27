package roomservice

import (
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path/filepath"

	"bookingrooms/internal/models"
	"github.com/nfnt/resize"
)

func (s service) SaveRoomImage(roomImage models.RoomsImages) error {
	return s.repository.CreateRoomImage(roomImage)
}

func (s service) readFileRoomImage(file *multipart.FileHeader, id int) (models.RoomsImages, error) {
	filePath := filepath.Join(filepath.Join("assets/image/", filepath.Base(file.Filename)))
	dst, err := os.Create(filePath)
	if err != nil {
		return models.RoomsImages{}, err
	}
	defer dst.Close()

	if file.Size >= 1000000 {
		s.resizeFileImage(file, dst)
	}

	// get root path
	cwd, err := os.Getwd()
	if err != nil {
		return models.RoomsImages{}, err
	}

	roomsImage := models.RoomsImages{
		RoomID:    id,
		ImagePath: cwd + "/" + filePath,
	}
	return roomsImage, err
}

func (s service) resizeFileImage(file *multipart.FileHeader, dst *os.File) error {
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

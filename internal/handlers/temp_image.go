package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handlers) SaveTempImages(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	files := form.File["images"]
	for _, file := range files {
		tempImage, err := h.tempImageService.ReadFileImage(file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = h.tempImageService.SaveTempImage(tempImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusCreated, "created temp images success")
}

func (h Handlers) DeleteTempImage(c echo.Context) error {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.tempImageService.SoftDeletedAtTempImages(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "deleted temp image success")
}

func (h Handlers) DisplayImage(c echo.Context) error {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	path, err := h.tempImageService.DisplayImage(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.File(path)

	//// not use but can use c.file
	//imageByte, err := os.ReadFile(TempImages.Path)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, err.Error())
	//}
	//
	//contentType := fmt.Sprintf("image/%s", TempImages.Extension)
	//c.Response().Header().Set("Content-Type", contentType)
	//c.Response().Write(imageByte)
	//return c.JSON(http.StatusOK, "Display image success")
}

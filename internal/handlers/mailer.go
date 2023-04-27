package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

func (h Handlers) Mailer(c echo.Context) error {
	requestImageId := c.QueryParam("image_id")

	imageId, err := strconv.Atoi(requestImageId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	image, err := h.tempImageService.GetTempImage(imageId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	message := gomail.NewMessage()
	message.SetHeader("From", "service.inetcvm@gmail.com")
	message.SetHeader("To", "mymimoza9@gmail.com", "uracha.moji@gmail.com")
	message.SetHeader("Subject", "Hello! Uracha")
	message.SetBody("text/html", "ทดสอบการส่ง Email ด้วย Golang <br> สวัสดี Uracha!")
	//message.Attach(image.Path)
	message.Attach("/Users/uracha.p/GolandProjects/booking-rooms/assets/file/Cyber Security.pdf")
	message.Embed(image.Path)
	err = h.mailerService.Send(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "send message success")
}

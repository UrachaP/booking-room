package handlers

import (
	bookingservice "bookingrooms/internal/services/booking"
	mailerservice "bookingrooms/internal/services/mailer"
	productservice "bookingrooms/internal/services/product"
	roomservice "bookingrooms/internal/services/room"
	tempimageservice "bookingrooms/internal/services/temp_image"
	userservice "bookingrooms/internal/services/user"
)

type Handlers struct {
	userService      userservice.UserService
	roomService      roomservice.RoomService
	tempImageService tempimageservice.TempImageService
	bookingService   bookingservice.BookingService
	mailerService    mailerservice.MailerService
	productService   productservice.ProductService
}

func NewHandlers(userService userservice.UserService, roomService roomservice.RoomService, tempImageService tempimageservice.TempImageService, bookingService bookingservice.BookingService, mailerService mailerservice.MailerService, productService productservice.ProductService) *Handlers {
	return &Handlers{
		userService:      userService,
		roomService:      roomService,
		tempImageService: tempImageService,
		bookingService:   bookingService,
		mailerService:    mailerService,
		productService:   productService,
	}
}

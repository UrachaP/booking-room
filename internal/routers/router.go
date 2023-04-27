package router

import (
	"bookingrooms/internal/handlers"
	"github.com/labstack/echo/v4"
)

type router struct {
	handlers *handlers.Handlers
}

func NewRouter(handlers *handlers.Handlers) router {
	return router{handlers: handlers}
}

func (r *router) MailerRouting(e *echo.Echo) {
	e.POST("/mailer", r.handlers.Mailer)
}

func (r *router) UserRouting(e *echo.Echo) {
	e.POST("/register", r.handlers.Register)
	e.POST("/login", r.handlers.Login)

	e.GET("/login/history", r.handlers.GetLoginHistory)
	e.PUT("/revoked/:id_token", r.handlers.RevokedAccessToken)

	userRoute := e.Group("/user")
	userRoute.GET("/:id", r.handlers.GetUser, JwtBasicMiddleware(), r.revokedMiddleware)
	userRoute.GET("", r.handlers.GetUserList)
	userRoute.PUT("/:id", r.handlers.UpdateUser, JwtBasicMiddleware())
	userRoute.PUT("/grade/:id", r.handlers.UpdateUserGrade, JwtBasicMiddleware())
	userRoute.DELETE("/:id", r.handlers.DeleteUser)
	userRoute.GET("/bookings/preload", r.handlers.PreloadUserBookings)
	userRoute.POST("/booking", r.handlers.CreateUserWithBooking)
}

func (r *router) RoomRouting(e *echo.Echo) {
	roomRoute := e.Group("/room")
	roomRoute.POST("", r.handlers.SaveRoom)
	roomRoute.GET("/:id", r.handlers.GetRoom)
	roomRoute.GET("", r.handlers.GetRoomList)
	roomRoute.PUT("/:id", r.handlers.UpdateRoom)
	roomRoute.DELETE("/:id", r.handlers.DeleteRoom)
	roomRoute.POST("/image/:id", r.handlers.SaveRoomImage)
}

func (r *router) BookingRouting(e *echo.Echo) {
	bookingRoute := e.Group("/booking")
	bookingRoute.POST("", r.handlers.SaveBooking)
	bookingRoute.GET("/:id", r.handlers.GetBooking)
	bookingRoute.GET("", r.handlers.GetBookingList)
	bookingRoute.DELETE("/:id", r.handlers.DeleteBooking)
	bookingRoute.PUT("/:id", r.handlers.UpdateBooking)
	bookingRoute.GET("/filter", r.handlers.GetBookingFilter)
}

func (r *router) TempImageRouting(e *echo.Echo) {
	TempImagesRoute := e.Group("/image")
	TempImagesRoute.POST("", r.handlers.SaveTempImages)
	TempImagesRoute.DELETE("/:id", r.handlers.DeleteTempImage)
	TempImagesRoute.GET("/:id", r.handlers.DisplayImage)
}

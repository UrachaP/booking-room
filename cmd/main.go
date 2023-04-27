package main

import (
	"log"

	"bookingrooms/config"
	"bookingrooms/internal/handlers"
	"bookingrooms/internal/mq"
	bookingrepository "bookingrooms/internal/repositories/booking"
	productrepository "bookingrooms/internal/repositories/product"
	roomrepository "bookingrooms/internal/repositories/room"
	tempimagerepository "bookingrooms/internal/repositories/temp_image"
	userrepository "bookingrooms/internal/repositories/user"
	router "bookingrooms/internal/routers"
	bookingservice "bookingrooms/internal/services/booking"
	mailerservice "bookingrooms/internal/services/mailer"
	productservice "bookingrooms/internal/services/product"
	roomservice "bookingrooms/internal/services/room"
	tempimageservice "bookingrooms/internal/services/temp_image"
	userservice "bookingrooms/internal/services/user"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	//open database
	db := config.InitDatabase()
	sqlDb, _ := db.DB()
	defer func() {
		err := sqlDb.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	//validate
	handlerValidate := config.Handlers{DB: db}
	e.Validator = handlerValidate.InitValidate()

	//mailer
	mailer := config.InitMailer()

	//open redis
	redisDb := config.InitRedis()
	defer func() {
		err := redisDb.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	//set init api
	userRepository := userrepository.NewUserRepository(db, redisDb)
	bookingRepository := bookingrepository.NewBookingRepository(db)
	roomRepository := roomrepository.NewRoomRepository(db)
	tempImageRepository := tempimagerepository.NewTempImageRepository(db)
	productRepository := productrepository.NewProductRepository(db)

	userService := userservice.NewUserService(userRepository)
	bookingService := bookingservice.NewBookingService(bookingRepository)
	roomService := roomservice.NewRoomService(roomRepository)
	tempImageService := tempimageservice.NewTempImageService(tempImageRepository)
	mailerService := mailerservice.NewMailerService(mailer)
	productService := productservice.NewProductService(productRepository)

	handler := handlers.NewHandlers(userService, roomService, tempImageService, bookingService, mailerService, productService)
	//handlerUser := handlers.NewHandlers(userService, roomService)

	api := router.NewRouter(handler)

	//call router
	api.UserRouting(e)
	api.MailerRouting(e)

	//test rabbitmq
	e.GET("/mq/send", mq.SendRabbitMQ)
	e.GET("/mq/receive", mq.ReceiveRabbitMQ)
	e.GET("/cut_stock", handler.BuyProduct)

	//cron jobs
	handler.RunCronJobs()

	e.Logger.Fatal(e.Start(":8080"))

}

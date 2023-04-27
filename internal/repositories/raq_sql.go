package repositories

import (
	"bookingrooms/internal/models"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) readHistoryBookingRooms() []models.HistoryBookingRooms {
	var historyBookingRooms []models.HistoryBookingRooms
	r.DB.Raw("SELECT CONCAT(COALESCE(name_prefix,'') , first_name , ' ' , last_name) AS full_name,room_name,maximum_person,amount_person,start_date,end_date\nFROM ((users\nINNER JOIN bookings ON users.id = bookings.users_id)\nINNER JOIN rooms ON rooms.id = bookings.rooms_id)\nWHERE start_date >= '2021-09-02 00:00:00' and start_date <= '2021-09-02 23:59:59'\nOR end_date >= '2021-09-02 00:00:00' and end_date <= '2021-09-02 23:59:59';").Scan(&historyBookingRooms)
	return historyBookingRooms
}

func (r Repository) readAmountOfBookingRooms() []models.AmountOfBookingRooms {
	var amountOfBookingRooms []models.AmountOfBookingRooms
	r.DB.Raw("SELECT room_name,count(bookings.id) AS COUNT\nFROM rooms \nLEFT JOIN bookings ON rooms.id = bookings.rooms_id\nGROUP BY rooms.id").Scan(&amountOfBookingRooms)
	return amountOfBookingRooms
}

func (r Repository) readAmountOfBookingUsers() []models.AmountOfBookingUsers {
	var amountOfBookingUsers []models.AmountOfBookingUsers
	r.DB.Raw("SELECT CONCAT(COALESCE(name_prefix,''),first_name,' ',last_name) AS full_name,COUNT(bookings.id) AS bookings_count\nFROM users\nLEFT JOIN bookings ON users.id = bookings.users_id\nGROUP BY users.id;").Scan(&amountOfBookingUsers)
	return amountOfBookingUsers
}

func (r Repository) printQuerySQL() {
	fmt.Println(r.readHistoryBookingRooms())
	fmt.Println(r.readAmountOfBookingUsers())
	fmt.Println(r.readAmountOfBookingRooms())
}

func (r Repository) ormReadHistoryBookingRooms() []models.HistoryBookingRooms {
	var historyBookingRooms []models.HistoryBookingRooms
	r.DB.
		//Select("CONCAT(COALESCE(name_prefix,'') , first_name , ' ' , last_name) AS full_name", "room_name", "maximum_person", "amount_person", "start_date", "end_date").
		Select([]string{
			"CONCAT(COALESCE(name_prefix,'') , first_name , ' ' , last_name) AS full_name",
			"room_name",
			"maximum_person",
			"amount_person",
			"start_date",
			"end_date",
		}).
		Joins("INNER JOIN bookings ON users.id = bookings.users_id").
		Joins("INNER JOIN rooms ON rooms.id = bookings.rooms_id").
		//Where("start_date >= '2021-09-02 00:00:00' and start_date <= '2021-09-02 23:59:59' OR end_date >= '2021-09-02 00:00:00' and end_date <= '2021-09-02 23:59:59'").
		Where(r.DB.
			Where("start_date >= '2021-09-02 00:00:00'").
			Where("start_date <= '2021-09-02 23:59:59'")).
		Or(r.DB.
			Where("end_date >= '2021-09-02 00:00:00'").
			Where("end_date <= '2021-09-02 23:59:59'")).
		Find(&historyBookingRooms)
	return historyBookingRooms
}

func (r Repository) ormReadAmountOfBookingRooms() []models.AmountOfBookingRooms {
	var amountOfBookingRooms []models.AmountOfBookingRooms
	r.DB.
		//Select("room_name", "count(bookings.id) as count").
		Select([]string{
			"room_name",
			"count(bookings.id) as count",
		}).
		Joins("left join bookings ON rooms.id = bookings.rooms_id").
		Group("rooms.id").
		Find(&amountOfBookingRooms)
	return amountOfBookingRooms
}

func (r Repository) ormReadAmountOfBookingUsers() []models.AmountOfBookingUsers {
	var amountOfBookingUsers []models.AmountOfBookingUsers
	//db.Select("CONCAT(COALESCE(name_prefix,''),first_name,' ',last_name) AS full_name", "count(bookings.id) AS bookings_count").
	r.DB.
		Select([]string{
			"CONCAT(COALESCE(name_prefix,''),first_name,' ',last_name) AS full_name",
			"count(bookings.id) AS bookings_count",
		}).
		Joins("LEFT JOIN bookings ON users.id = bookings.users_id").
		Group("users.id").
		Find(&amountOfBookingUsers)
	return amountOfBookingUsers
}

func (r Repository) ormPrintQuerySQL() {
	log.Println("History Booking Rooms")
	for _, history := range r.ormReadHistoryBookingRooms() {
		b, err := json.Marshal(history)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}
	fmt.Println("===========================================")
	fmt.Println("Amount Of Booking Users")
	for _, users := range r.ormReadAmountOfBookingUsers() {
		b, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}
	fmt.Println("===========================================")
	fmt.Println("Amount Of Booking Rooms")
	for _, rooms := range r.ormReadAmountOfBookingRooms() {
		b, err := json.Marshal(rooms)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}
}

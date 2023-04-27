package roomrepository

import (
	"bookingrooms/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return repository{db: db}
}

type RoomRepository interface {
	GetRoomById(id int) (models.Rooms, error)
	CreateRoom(room models.Rooms) error
	GetCountRoomById(id int) (int, error)
	GetRooms() ([]models.Rooms, error)
	UpdateRoom(room models.Rooms) error
	DeleteRoom(room *models.Rooms) error
	CreateRoomImage(roomImage models.RoomsImages) error
}

func (r repository) GetCountRoomById(id int) (int, error) {
	var count int
	return count, r.db.
		Model(&models.Rooms{}).
		Select("COUNT(*)").
		Where("id = ?", id).
		First(&count).
		Error
}

func (r repository) CreateRoom(room models.Rooms) error {
	return r.db.Create(&room).Error
}

func (r repository) GetRoomById(id int) (models.Rooms, error) {
	var room models.Rooms
	return room, r.db.Where("id = ?", id).First(&room).Error
}

func (r repository) GetRooms() ([]models.Rooms, error) {
	var rooms []models.Rooms
	return rooms, r.db.Find(&rooms).Error
}

func (r repository) UpdateRoom(room models.Rooms) error {
	return r.db.Updates(&room).Error
}

func (r repository) DeleteRoom(room *models.Rooms) error {
	return r.db.Unscoped().Delete(&room).Error
}

package models

type Rooms struct {
	ID            int    `json:"id"`
	RoomName      string `json:"room_name" validate:"required"`
	MaximumPerson int    `json:"maximum_person" validate:"required"`
	Name          string `json:"name" validate:"required"`
}

type RoomsImages struct {
	ID        int    `json:"id"`
	RoomID    int    `json:"room_id"`
	ImagePath string `json:"image_path"`
	ImageId   int    `json:"image_id"`
}

func (AmountOfBookingRooms) TableName() string {
	return "rooms"
}

package models

type TempImages struct {
	Model
	Path      string `json:"path"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Temp      string `json:"temp" gorm:"default:1"`
}

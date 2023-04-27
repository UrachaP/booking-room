package models

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type RequestProduct struct {
	ID     int `json:"id"`
	Amount int `json:"amount"`
}

func (Product) TableName() string {
	return "product"
}

func (RequestProduct) TableName() string {
	return "product"
}

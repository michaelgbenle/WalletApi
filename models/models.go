package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `json:"name"`
	Balance uint   `json:"balance"`
}

type Transaction struct {
	gorm.Model
	CustomerId string `gorm:"foreignKey" json:"customer_id"`
	Type       string
	Success    bool `json:"success"`
}

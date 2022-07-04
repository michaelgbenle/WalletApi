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
	Success    bool   `json:"success"`
}

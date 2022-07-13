package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name       string `json:"name" binding:"required"`
	AccountNos string `gorm:"unique",json:"accountNos" binding:"required,email"`
	Balance    uint   `json:"balance"`
}

type Transaction struct {
	gorm.Model
	CustomerId uint   `gorm:"foreignKey" json:"customer_id"`
	AccountNos string `gorm:"foreignKey" json:"accountNos"`
	Type       string `json:"type"`
	Success    bool   `json:"success"`
}

type Money struct {
	AccountNos string
	Amount     uint
}

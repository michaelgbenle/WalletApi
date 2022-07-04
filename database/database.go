package database

import "gorm.io/gorm"

type PostgresDb struct {
	DB *gorm.DB
}

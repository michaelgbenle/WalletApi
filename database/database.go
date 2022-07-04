package database

import (
	"fmt"
	"github.com/michaelgbenle/WalletApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresDb struct {
	DB *gorm.DB
}

func (pdb *PostgresDb) SetupDb(host, user, password, dbName, port string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	pdb.DB = db

	dberr := pdb.DB.AutoMigrate(&models.Customer{})
	if dberr != nil {
		log.Fatal(dberr)
	}
	return nil
}

func (pdb *PostgresDb) wallet(id string) (*models.Wallet, error) {
	var customer *models.Wallet
	if err := pdb.DB.Where("id=?", id).First(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

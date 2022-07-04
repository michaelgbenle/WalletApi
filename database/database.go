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

	dberr := pdb.DB.AutoMigrate(&models.Customer{}, models.Transaction{})
	if dberr != nil {
		log.Fatal(dberr)
	}
	return nil
}

func (pdb *PostgresDb) Getcustomer(id string) (*models.Customer, error) {
	var customer *models.Customer
	if err := pdb.DB.Where("id=?", id).First(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (pdb *PostgresDb) Addcustomer(user *models.Customer) error {
	return nil
}

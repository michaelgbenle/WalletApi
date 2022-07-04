package database

import (
	"github.com/joho/godotenv"
	"github.com/michaelgbenle/WalletApi/models"
	"log"
	"os"
)

type DB interface {
	Getcustomer(id string) (*models.Customer, error)
	Addcustomer(customer *models.Customer) error
	Creditwallet(money *models.Money) error
	Debitwallet(money models.Money) error
}

type DbParameters struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func InitializeDbParameters() DbParameters {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return DbParameters{
		Host:     host,
		User:     user,
		Password: password,
		DbName:   dbName,
		Port:     port,
	}
}

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

func (pdb *PostgresDb) Getcustomer(accountNos string) (*models.Customer, error) {
	var customer *models.Customer
	if err := pdb.DB.Where("account_nos=?", accountNos).First(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (pdb *PostgresDb) Creditwallet(money *models.Money) (*models.Transaction, error) {
	accountNos, amount := money.AccountNos, money.Amount
	user, _ := pdb.Getcustomer(accountNos)
	transaction := &models.Transaction{
		CustomerId: user.ID,
		AccountNos: user.AccountNos,
		Type:       "credit",
		Success:    false,
	}
	if err := pdb.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}

	if err := pdb.DB.Model(user).Where("account_nos=?", accountNos).
		Update("balance", user.Balance+amount).
		Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := pdb.DB.Model(transaction).Where("Customer_id=?", user.ID).
		Update("success", true).
		Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (pdb *PostgresDb) CreateTransaction(transaction *models.Transaction) {
	pdb.DB.Create(&transaction)
}

func (pdb *PostgresDb) Debitwallet(money *models.Money) (*models.Transaction, error) {
	accountNos, amount := money.AccountNos, money.Amount
	user, _ := pdb.Getcustomer(accountNos)
	transaction := &models.Transaction{
		CustomerId: user.ID,
		AccountNos: user.AccountNos,
		Type:       "debit",
		Success:    false,
	}
	if err := pdb.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}

	if err := pdb.DB.Model(user).Where("account_nos=?", accountNos).
		Update("balance", user.Balance-amount).
		Error; err != nil {
		return nil, err
	}
	if err := pdb.DB.Model(transaction).Where("Customer_id=?", user.ID).
		Update("success", true).
		Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
func (pdb *PostgresDb) Gettransaction(accountNos string) (*[]models.Transaction, error) {
	var transactions *[]models.Transaction
	if err := pdb.DB.Where("account_nos=?", accountNos).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (pdb *PostgresDb) Addcustomer(customer models.Customer) error {
	customer.Balance = 0
	if err := pdb.DB.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

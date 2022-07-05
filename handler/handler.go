package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/WalletApi/database"
	"github.com/michaelgbenle/WalletApi/models"
	"log"
	"net/http"
)

type Handler struct {
	DB database.DB
}

func (h *Handler) GetCustomer(c *gin.Context) {
	accountNos := c.Query("accountNos")
	if len(accountNos) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account number should be 10 digits"})
		return
	}
	customer, err := h.DB.Getcustomer(accountNos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": customer,
	})
}

func (h *Handler) CreditWallet(c *gin.Context) {
	credit := &models.Money{}
	if err := c.ShouldBindJSON(credit); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}
	if len(credit.AccountNos) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account number should be 10 digits"})
	}

	transaction, CreditErr := h.DB.Creditwallet(credit)
	if CreditErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to credit wallet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "wallet credited successfully",
		"transaction": transaction,
	})

}
func (h *Handler) DebitWallet(c *gin.Context) {
	debit := &models.Money{}
	if err := c.ShouldBindJSON(debit); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}

	customer, _ := h.DB.Getcustomer(debit.AccountNos)
	if customer.Balance < debit.Amount {
		transaction := &models.Transaction{
			CustomerId: customer.ID,
			AccountNos: customer.AccountNos,
			Type:       "debit",
			Success:    false,
		}
		h.DB.CreateTransaction(transaction)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "insufficient funds"})
		return
	}

	transaction, debitErr := h.DB.Debitwallet(debit)
	if debitErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to debit wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "wallet debited successfully",
		"transaction": transaction,
	})

}
func (h *Handler) GetTransaction(c *gin.Context) {
	id := c.Query("id")
	transactions, err := h.DB.Gettransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "transaction not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": transactions,
	})

}
func (h *Handler) AddCustomer(c *gin.Context) {
	var customer *models.Customer
	log.Println(customer)
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}
	fmt.Println("test 2")
	fmt.Println(customer)
	if user, userErr := h.DB.Getcustomer(customer.AccountNos); userErr == nil {
		fmt.Println("this is ", user)
		fmt.Println("our ", userErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer exists"})
		return
	}

	if CreateErr := h.DB.Addcustomer(*customer); CreateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to create customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully"})

}

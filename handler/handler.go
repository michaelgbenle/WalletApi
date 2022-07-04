package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/WalletApi/database"
	"github.com/michaelgbenle/WalletApi/models"
	"net/http"
)

type Handler struct {
	DB database.DB
}

func (h *Handler) GetCustomer(c *gin.Context) {
	id := c.Query("id")
	customer, err := h.DB.Getcustomer(id)
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
	if err := c.ShouldBindJSON(credit).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}
	transaction, CreditErr := h.DB.Creditwallet(credit)
	if CreditErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to credit wallet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "wallet credited successfully"})

}
func (h *Handler) DebitWallet(c *gin.Context) {
	debit := &models.Money{}
	if err := c.ShouldBindJSON(debit).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}
	if transaction, DebitErr := h.DB.Debitwallet(debit).Error; DebitErr != nil {
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
	transaction, err := h.DB.Gettransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "transaction not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": transaction,
	})

}
func (h *Handler) AddCustomer(c *gin.Context) {
	customer := &models.Customer{}
	if err := c.ShouldBindJSON(customer).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}
	if CreateErr := h.DB.Addcustomer(customer).Error; CreateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to create customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully"})

}

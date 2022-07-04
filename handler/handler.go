package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/WalletApi/database"
	"net/http"
)

type Handler struct {
	DB database.DB
}

func (h *Handler) GetWallets(c *gin.Context) {
	id := c.Query("id")
	customer, err := h.DB.Wallet(id)
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
func (h *Handler) GetTransaction(c *gin.Context) {

}

func (h *Handler) CreditWallet(c *gin.Context) {
	id := c.Query("id")

}
func (h *Handler) DebitWallet(c *gin.Context) {

}

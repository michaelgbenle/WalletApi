package router

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/WalletApi/handler"
	"os"
)

func SetupRouter(h *handler.Handler) (*gin.Engine, string) {
	router := gin.Default()
	router.GET("/customer", h.GetCustomer)
	router.GET("/transaction", h.GetTransaction)
	router.PATCH("/credit", h.CreditWallet)
	router.PATCH("/return", h.DebitWallet)

	port := os.Getenv("PORT")

	return router, port
}

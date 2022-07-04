package server

import (
	"github.com/michaelgbenle/WalletApi/database"
	"github.com/michaelgbenle/WalletApi/handler"
	"github.com/michaelgbenle/WalletApi/router"
	"log"
)

func Start() error {
	values := database.InitializeDbParameters()
	var PDB = new(database.PostgresDb)
	h := &handler.Handler{DB: PDB}

	err := PDB.SetupDb(values.Host, values.User, values.Password, values.DbName, values.Port)
	if err != nil {
		log.Fatal(err)
	}
	routes, port := router.SetupRouter(h)
	err = routes.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

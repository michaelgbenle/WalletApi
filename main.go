package main

import (
	"github.com/michaelgbenle/WalletApi/server"
	"log"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

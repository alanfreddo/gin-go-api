package main

import (
	"github.com/alanfreddo/gin-go-api/database"
	"github.com/alanfreddo/gin-go-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()

}

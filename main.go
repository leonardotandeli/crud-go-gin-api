package main

import (
	"gogin/database"
	"gogin/routes"
)

func main() {
	database.ConectaComOBancoDeDados()

	routes.HandleRequests()

}

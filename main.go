package main

import (
	"fmt"

	"github.com/matheuspolitano/my-api-go/database"
	"github.com/matheuspolitano/my-api-go/models"
	"github.com/matheuspolitano/my-api-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConnectWithDB()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}

package main

import (
	"fmt"

	"github.com/lcardelli/curso/routes"
)

func main() {
	fmt.Println("Servidor rodado na porta 8080")
	routes.HandleRequests()
}

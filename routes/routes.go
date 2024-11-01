package routes

import (
	"log"
	"net/http"

	"github.com/lcardelli/curso/handlers"
)

func HandleRequests() {
	http.HandleFunc("/", handlers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"net/http"

	"github.com/RaphaSalomao/alura-webapp-go/router"
)

func main() {
	router.HandleRoutes()
	http.ListenAndServe(":8000", nil)
}

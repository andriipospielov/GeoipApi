package main

import (
	"github.com/andriipospielov/GeoipApi/controller"
	"github.com/andriipospielov/GeoipApi/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server")
	router := mux.NewRouter()
	ipController := controller.NewCrudController()
	router.Handle("/{source}/{needle}", middleware.AuthMiddleware(http.HandlerFunc(ipController.Index)))
	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

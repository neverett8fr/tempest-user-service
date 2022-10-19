package main

import (
	"fmt"
	"log"
	"tempest-user-service/cmd"
	application "tempest-user-service/pkg/application/service"
	"tempest-user-service/pkg/config"

	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes() *mux.Router {
	r := mux.NewRouter()
	application.NewUserInformation(r)

	return r
}

// Initiate web server
func main() {

	conf := config.Initialise()

	fmt.Println("service started")
	router := getRoutes()

	err := cmd.StartServer(*conf, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
	}

}

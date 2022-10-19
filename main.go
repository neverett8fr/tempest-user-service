package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"tempest-user-service/cmd"
	"tempest-user-service/pkg/config"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

// Initiate web server
func main() {

	conf := config.Initialise()

	fmt.Println("service started")
	router := router()

	err := cmd.StartServer(*conf, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
	}

}

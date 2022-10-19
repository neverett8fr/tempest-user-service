package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"tempest-user-service/pkg/config"
	"time"

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
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("0.0.0.0.:%v", conf.Service.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started on port: %v", conf.Service.Port)

	log.Fatal(srv.ListenAndServe())
}

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"tempest-user-service/pkg/config"
	"time"

	"github.com/gorilla/mux"
)

func StartServer(conf config.Config, router *mux.Router) error {
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%v:%v", conf.Service.Host, conf.Service.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started on port: %v", conf.Service.Port)

	log.Fatal(srv.ListenAndServe())
	return nil
}

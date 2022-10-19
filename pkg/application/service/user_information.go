package application

import (
	"net/http"
	"tempest-user-service/pkg/application/entities"

	"github.com/gorilla/mux"
)

func NewUserInformation(r *mux.Router) {
	r.HandleFunc("/test/{text}", testHandler).Methods("GET")
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := entities.NewResponse(text)

	writeReponse(w, r, body)
}

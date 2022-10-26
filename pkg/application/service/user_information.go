package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tempest-user-service/pkg/entities"
	"tempest-user-service/pkg/infra/db"

	"github.com/gorilla/mux"
)

func NewUserInformation(r *mux.Router, conn *sql.DB) {
	DBConn = db.NewDBConnFromExisting(conn)

	r.HandleFunc("/test/{text}", testHandler).Methods("GET")
	r.HandleFunc("/user", createUserHandler).Methods("POST")
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := NewResponse(fmt.Sprintf("test: %v", text))

	writeReponse(w, body)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	bodyIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}
	userInformation := newUserIn{}
	json.Unmarshal(bodyIn, &userInformation)

	user, err := entities.NewUser(userInformation.Username, userInformation.Password)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	err = DBConn.CreateUser(user)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	body := NewResponse(fmt.Sprintf("user created with username %v", userInformation.Username), err)

	writeReponse(w, body)
}

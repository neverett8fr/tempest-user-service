package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"tempest-user-service/pkg/infra/db"

	"github.com/gorilla/mux"
)

func NewUserInformation(r *mux.Router, conn *sql.DB) {
	DBConn = db.NewDBConnFromExisting(conn)

	r.HandleFunc("/test/{text}", testHandler).Methods(http.MethodGet)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := NewResponse(fmt.Sprintf("test: %v", text))

	writeReponse(w, body)
}

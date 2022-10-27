package service

import (
	"encoding/json"
	"log"
	"net/http"
	"tempest-user-service/pkg/infra/db"
)

var (
	DBConn *db.DBConn
)

type Response struct {
	Data   interface{} `json:"data"`
	Errors []error     `json:"errors"`
}

func NewResponse(data interface{}, err ...error) *Response {

	return &Response{
		Data:   data,
		Errors: err,
	}
}

func writeReponse(w http.ResponseWriter, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
		return
	}
}

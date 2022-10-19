package application

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeReponse(w http.ResponseWriter, r *http.Request, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}

	w.Write(reponseBody)
}

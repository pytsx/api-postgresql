package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pytsx/api-postgresql/model"
)

func List(res http.ResponseWriter, req *http.Request) {
	todos, err := model.GetAll()

	if err != nil {
		log.Printf("erro ao obter registros: %v", err)
	}

	
	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(todos)

}
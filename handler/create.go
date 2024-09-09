package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pytsx/api-postgresql/model"
)

func Create(res http.ResponseWriter, req *http.Request) {
	var todo model.Todo

	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Printf("error ao fazer o decode do json: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	id, err := model.Insert(todo)
	
	var resp map[string]any 

	if err != nil {
		resp = map[string]any {
			"Error": true, 
			"Message": fmt.Sprintf("Erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any {
			"Error": false,
			"Message": fmt.Sprintf("todo inserido com sucesso! ID: %d", id),
		}
	}


	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
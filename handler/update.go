package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pytsx/api-postgresql/model"
)

func Update(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Printf("error ao fazer parse do id: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	var todo model.Todo

	err = json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Printf("error ao fazer o decode do json: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	rows, err := model.Update(int64(id), todo)
	if err != nil {
		log.Printf("error ao atualiar registro: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	if rows > 1 {
		log.Printf("foram atualizados %d registros", rows)
	}


	resp := map[string]any {
		"Error": false,
		"Message": "dados atualizados com sucesso!",
	}

	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
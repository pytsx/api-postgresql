package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pytsx/api-postgresql/model"
)

func Delete(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Printf("error ao fazer parse do id: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	rows, err := model.Delete(int64(id))
	if err != nil {
		log.Printf("error ao remover registro: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	if rows > 1 {
		log.Printf("foram removidos %d registros", rows)
	}


	resp := map[string]any {
		"Error": false,
		"Message": "registro removido com sucesso!",
	}

	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
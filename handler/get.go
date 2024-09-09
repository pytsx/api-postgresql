package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pytsx/api-postgresql/model"
)

func Get(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Printf("error ao fazer parse do id: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	todo, err := model.Get(int64(id))
	if err != nil {
		log.Printf("error ao resgatar registro %d: %v", id, err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}
	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(todo)
}
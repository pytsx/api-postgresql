package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pytsx/api-postgresql/config"
	"github.com/pytsx/api-postgresql/handler"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)
	router.Delete("/{id}", handler.Delete)
	router.Get("/", handler.List)
	router.Get("/{id}", handler.Get)

	http.ListenAndServe(":"+config.GetServerPort(), router)
}
package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, render.M{"hello": "world"})
	})

	http.ListenAndServe(":8080", r)
}

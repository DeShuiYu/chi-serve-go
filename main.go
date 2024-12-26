package main

import (
	"github.com/go-chi/chi/v5"
	"net"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server := &http.Server{
		Handler: r,
	}
	l, _ := net.Listen("tcp", ":8080")
	server.Serve(l)
}

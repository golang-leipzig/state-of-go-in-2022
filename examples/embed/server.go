package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed assets
var assets embed.FS

func main() {
	server := http.Server{
		Addr:    "localhost:11111",
		Handler: router(),
	}

	log.Printf("Accepting connections on http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func router() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Try /assets/leipzig-gopher/leipzig-gopher.png")
	}))
	router.Handle("/assets/", http.FileServer(http.FS(assets)))

	return router
}

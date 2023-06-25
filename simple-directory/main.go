package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed htmx-static
var htmxAssets embed.FS

func main() {

	mux := http.NewServeMux()
	mux.Handle("/htmx-static/", http.FileServer(http.FS(htmxAssets)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

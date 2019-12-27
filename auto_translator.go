package main

import (
	"fmt"
	"github.com/nickpasko/auto_translator/handlers"
	"net/http"
)

func main() {
	addr := ":8081"

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/index/", handlers.IndexHandler)
	http.HandleFunc("/search/", handlers.SearchHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Listening on addr: %s\n", addr)
	http.ListenAndServe(addr, nil)
}

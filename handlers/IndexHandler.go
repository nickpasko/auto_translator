package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, "Frag! Template not found!")
		log.Println("Frag! Template not found!")
		return
	}

	index.Execute(w, nil)
}

// call by proper IP, "localhost" doesn't really work

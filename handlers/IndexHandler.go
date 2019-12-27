package handlers

import (
	"net/http"
	"fmt"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, "Frag! Template not found!")
		return
	}

	index.Execute(w, nil)
}

//192.168.88.253
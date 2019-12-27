package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/nickpasko/auto_translator/auto_model"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	search, err := template.ParseFiles("templates/search.html")
	if err != nil {
		fmt.Fprint(w, "Frag! Template not found!")
		return
	}

	search.Execute(w, nil)
}

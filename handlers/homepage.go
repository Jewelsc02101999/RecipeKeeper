package handlers

import (
	"codingforwomen/data"
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/homePage.html")
	if err != nil {
		log.Fatal(err, "ERROR: problem with home file path")
	}
	tmpl.Execute(w, data.AllRecipes)

}

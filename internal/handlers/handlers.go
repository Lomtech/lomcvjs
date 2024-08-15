package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		log.Fatalf("Failed to determine project root: %v", err)
	}
	tmplPath := filepath.Join(projectRoot, "internal/templates/index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Setze die Pfade zu den Ressourcen
	imagesPath := "templates/images"
	cssPath := "templates/css"

	// Setze die Handler für die Routen
	http.HandleFunc("/", HomeHandler)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(imagesPath))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(cssPath))))

	// Starte den Server
	fmt.Println("Server läuft auf http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

// HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

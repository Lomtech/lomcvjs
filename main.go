package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/lomtech/asrudi/internal/handlers"
)

func main() {

	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		log.Fatalf("Failed to determine project root: %v", err)
	}

	imagesPath := filepath.Join(projectRoot, "internal/templates/images")

	http.HandleFunc("/", handlers.HomeHandler)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(imagesPath))))

	fmt.Println("Server läuft auf http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

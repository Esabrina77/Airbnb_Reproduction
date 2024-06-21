package main

import (
	"AIRBNB/routeur"
	initTemplate "AIRBNB/templates"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback au port 8080 si la variable d'environnement PORT n'est pas définie
	}

	fmt.Println("server is running...")
	fmt.Println("")
	fmt.Printf("CLICK HERE to OPEN PAGE---> http://localhost:%s/home \n", port)
	fmt.Println("")
	fmt.Println("TO STOP THE SERVER, PRESS 'ctrl+C'")

	initTemplate.InitTemplate()
	
	// Initialize your router and server
	router := routeur.InitServe()

	log.Fatal(http.ListenAndServe(":"+port, router))
}

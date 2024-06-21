package main

import (
    "AIRBNB/routeur"
    initTemplate "AIRBNB/templates"
    "fmt"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Println("server is running...")

    fmt.Println("")
    fmt.Printf("CLICK HERE to OPEN PAGE---> http://localhost:%s/home \n", port)
    fmt.Println("")
    fmt.Println("TO STOP THE SERVER, PRESS 'ctrl+C' ")
    initTemplate.InitTemplate()
    routeur.InitServe()
}

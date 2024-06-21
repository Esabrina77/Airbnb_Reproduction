package routeur

import (
    "log"
    "net/http"
    "os"
)

func InitServe() {
    fileServer := http.FileServer(http.Dir("CSS"))
    http.Handle("/CSS/", http.StripPrefix("/CSS/", fileServer))

    http.HandleFunc("/home", controller.HomeProduct)
    http.HandleFunc("/cabane", controller.CabaneProduct)
    http.HandleFunc("/wow", controller.WowProduct)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Fallback au port 8080 si la variable d'environnement PORT n'est pas définie
    }

    if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
        log.Fatal(err)
    }
}

package routeur

import (
	"AIRBNB/controller"
	"log"
	"net/http"
)

func InitServe() {
	fileServer := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fileServer))

	http.HandleFunc("/home", controller.DisplayProduct)

	if err := http.ListenAndServe(controller.Port, nil); err != nil {
		log.Fatal(err)
	}
}

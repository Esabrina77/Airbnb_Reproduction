package routeur

import (
	"AIRBNB/controller"
	"log"
	"net/http"
)

func InitServe() {
	fileServer := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fileServer))

	http.HandleFunc("/home", controller.HomeProduct)
	http.HandleFunc("/cabane", controller.CabaneProduct)
	http.HandleFunc("/wow", controller.WowProduct)
	if err := http.ListenAndServe(controller.Port, nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"AIRBNB/routeur"
	initTemplate "AIRBNB/templates"
	"fmt"
)

func main() {
	fmt.Println("server is running...")

	fmt.Println("")
	fmt.Print("CLICK HERE to OPEN PAGE---> http://localhost:8080/home \n")
	fmt.Println("")
	fmt.Println("TO STOP THE SERVER , PRESS  'ctrl+C' ")
	initTemplate.InitTemplate()
	routeur.InitServe()

}

package initTemplate

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplate() {

	temp, err := template.ParseGlob("templates/*.html")

	if err != nil {
		fmt.Printf("ERREUR SURVENUE LORS DE L OUVERTURE DES TEMPLATES: %v", err.Error())
		os.Exit(1)
	}
	Temp = temp
}

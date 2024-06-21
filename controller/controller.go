package controller

import (
	AirBnb_test "AIRBNB/Air"
	initTemplate "AIRBNB/templates"
	"net/http"
)

const Port = "0.0.0.0:8080"

func HomeProduct(w http.ResponseWriter, r *http.Request) {
	// Initialisation des données
	Screen := AirBnb_test.ProductData{
		Products: []AirBnb_test.Product{
			{Localisation: "Paris,France", Views: "City Views", Date: "DEC 6-11", Price: 615, Images: "../CSS/pictures/chambre.png"},
			{Localisation: "Marseille,France", Views: " 2 kilometers away", Date: "DEC 13-12", Price: 542, Images: "CSS/pictures/Salon.png"},
			{Localisation: "Toulouse,France", Views: "4 kilometers away", Date: "DEC 7-11", Price: 320, Images: "/CSS/pictures/Terasse.png"},
			{Localisation: "Bouffémont,Île-de-France,France", Views: "Professionnel", Date: "DEC 9-14", Price: 5049, Images: "../CSS/pictures/Bouffémont, Île-de-France,France.webp"},

			{Localisation: "Châteaufort,France", Views: " Particulier", Date: "DEC 3-8", Price: 316, Images: "CSS/pictures/Châteaufort,France.webp"},
			{Localisation: "Le Coudray-sur-Thelle, France", Views: "Particulier", Date: "MARS 1-6", Price: 124, Images: "/CSS/pictures/Le Coudray-sur-Thelle,France.webp"},
			{Localisation: "Le Vésinet,France", Views: "Professionnel", Date: "DEC 18-23", Price: 327, Images: "../CSS/pictures/Le Vésinet,France.webp"},
			{Localisation: "Maisons-Laffitte,France", Views: " Particulier", Date: "DEC 2-7", Price: 179, Images: "CSS/pictures/Maisons-Laffitte,France.webp"},
			{Localisation: "Montreuil-sur-Epte", Views: "Professionnel", Date: "DEC 10-15", Price: 257, Images: "/CSS/pictures/Montreuil-sur-Epte,France2.webp"},

			{Localisation: "Montreuil-sur-Epte", Views: "Professionnel", Date: "DEC 16-21", Price: 257, Images: "../CSS/pictures/Montreuil-sur-Epte.webp"},
			{Localisation: "Raray,France", Views: " Particulier", Date: "FEV 15-20", Price: 96, Images: "CSS/pictures/Raray,France.webp"},
			{Localisation: "Rennemoulin,France", Views: "Professionnel", Date: "DEC 5-10", Price: 424, Images: "/CSS/pictures/Rennemoulin,France.webp"},
			{Localisation: "Vaucresson,France", Views: "Particulier", Date: "DEC 5-10", Price: 409, Images: "../CSS/pictures/Vaucresson,France.webp"},
		},
	}
	// Réponse : envoie du template avec les données au client
	initTemplate.Temp.ExecuteTemplate(w, "home", Screen)
}

func CabaneProduct(w http.ResponseWriter, r *http.Request) {
	// Initialisation des données
	Screen := AirBnb_test.ProductData{
		Products: []AirBnb_test.Product{
			{Localisation: "Bouffémont,Île-de-France,France", Views: "Professionnel", Date: "DEC 9-14", Price: 5049, Images: "../CSS/pictures/Bouffémont, Île-de-France,France.webp"},

			{Localisation: "Châteaufort,France", Views: " Particulier", Date: "DEC 3-8", Price: 316, Images: "CSS/pictures/Châteaufort,France.webp"},
			{Localisation: "Le Coudray-sur-Thelle, France", Views: "Particulier", Date: "MARS 1-6", Price: 124, Images: "/CSS/pictures/Le Coudray-sur-Thelle,France.webp"},
			{Localisation: "Le Vésinet,France", Views: "Professionnel", Date: "DEC 18-23", Price: 327, Images: "../CSS/pictures/Le Vésinet,France.webp"},
			{Localisation: "Maisons-Laffitte,France", Views: " Particulier", Date: "DEC 2-7", Price: 179, Images: "CSS/pictures/Maisons-Laffitte,France.webp"},
			{Localisation: "Montreuil-sur-Epte", Views: "Professionnel", Date: "DEC 10-15", Price: 257, Images: "/CSS/pictures/Montreuil-sur-Epte,France2.webp"},

			{Localisation: "Montreuil-sur-Epte", Views: "Professionnel", Date: "DEC 16-21", Price: 257, Images: "../CSS/pictures/Montreuil-sur-Epte.webp"},
			{Localisation: "Raray,France", Views: " Particulier", Date: "FEV 15-20", Price: 96, Images: "CSS/pictures/Raray,France.webp"},
		},
	}
	// Réponse : envoie du template avec les données au client
	initTemplate.Temp.ExecuteTemplate(w, "cabane", Screen)
}

func WowProduct(w http.ResponseWriter, r *http.Request) {
	// Initialisation des données
	Screen := AirBnb_test.ProductData{
		Products: []AirBnb_test.Product{
			{Localisation: "Paris,France", Views: "City Views", Date: "DEC 6-11", Price: 615, Images: "../CSS/pictures/chambre.png"},
			{Localisation: "Marseille,France", Views: " 2 kilometers away", Date: "DEC 13-12", Price: 542, Images: "CSS/pictures/Salon.png"},
			{Localisation: "Toulouse,France", Views: "4 kilometers away", Date: "DEC 7-11", Price: 320, Images: "/CSS/pictures/Terasse.png"},
			{Localisation: "Bouffémont,Île-de-France,France", Views: "Professionnel", Date: "DEC 9-14", Price: 5049, Images: "../CSS/pictures/Bouffémont, Île-de-France,France.webp"},

			{Localisation: "Châteaufort,France", Views: " Particulier", Date: "DEC 3-8", Price: 316, Images: "CSS/pictures/Châteaufort,France.webp"},
			{Localisation: "Le Coudray-sur-Thelle, France", Views: "Particulier", Date: "MARS 1-6", Price: 124, Images: "/CSS/pictures/Le Coudray-sur-Thelle,France.webp"},
			{Localisation: "Le Vésinet,France", Views: "Professionnel", Date: "DEC 18-23", Price: 327, Images: "../CSS/pictures/Le Vésinet,France.webp"},
			{Localisation: "Rennemoulin,France", Views: "Professionnel", Date: "DEC 5-10", Price: 424, Images: "/CSS/pictures/Rennemoulin,France.webp"},
			{Localisation: "Vaucresson,France", Views: "Particulier", Date: "DEC 5-10", Price: 409, Images: "../CSS/pictures/Vaucresson,France.webp"},
		},
	}
	// Réponse : envoie du template avec les données au client
	initTemplate.Temp.ExecuteTemplate(w, "wow", Screen)
}

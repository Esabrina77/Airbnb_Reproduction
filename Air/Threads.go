package AirBnb_test

type Product struct {
	Localisation string
	Views        string
	Date         string
	Price        int
	Images       string
}
type ProductData struct {
	Products []Product
}

var Screen ProductData

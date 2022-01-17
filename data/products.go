package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

func getProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Esspresso",
		Description: "Esspresso Coffee",
		SKU:         "coffee-1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          2,
		Name:        "Cafe Latte",
		Description: "Latte with milk",
		SKU:         "coffee-2",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}

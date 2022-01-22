package data

import (
	"fmt"
	"time"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[0-9]+
	SKU       string `json:"sku" validate:"sku"`
	CreatedOn string `json:"_"`
	UpdatedOn string `json:"_"`
	DeletedOn string `json:"_"`
}

//Collection of Proucts
type Products []*Product

//Get all products
func GetProducts() Products {
	return productList
}

//Get single product
func GetProductById(id int) (*Product, error) {
	_, pos, err := findIndexByProductId(id)
	if err != nil {
		return nil, ErrProductNotFound
	}

	return productList[pos], nil
}

//Add product
func AddProduct(p *Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, p)
}

//Update product by ID
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findIndexByProductId(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

//Delete product by ID
func DeleteProduct(id int) error {
	_, pos, err := findIndexByProductId(id)
	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)

	return nil
}

func findIndexByProductId(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var ErrProductNotFound = fmt.Errorf("Product not Found!")

//Temporary product storage
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Esspresso",
		Description: "Esspresso Coffee",
		Price:       2.75,
		SKU:         "coffee-1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          2,
		Name:        "Cafe Latte",
		Description: "Latte with milk",
		Price:       2.00,
		SKU:         "coffee-2",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}

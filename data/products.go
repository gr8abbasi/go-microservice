package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"_"`
	UpdatedOn   string  `json:"_"`
	DeletedOn   string  `json:"_"`
}

//Collection of Proucts
type Products []*Product

//Encode ToJSON
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

//Decode FromJSON
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

// Validate Prouct
func (p *Product) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("sku", validateSku)
	return validator.Struct(p)
}

func validateSku(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`[a-z]+-[0-9]+`)
	matches := regex.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

//Get all products
func GetProducts() Products {
	return productList
}

//Add product
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

//get product next ID
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

//Update product
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProductById(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func findProductById(id int) (*Product, int, error) {
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

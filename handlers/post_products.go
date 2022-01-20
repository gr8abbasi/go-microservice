package handlers

import (
	"net/http"

	"github.com/gr8abbasi/go-microservice/data"
)

// swagger:route POST /products products createProduct
// Create a new product
// responses:
// 	200: productsResponse
//	422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

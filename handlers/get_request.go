package handlers

import (
	"net/http"

	"github.com/gr8abbasi/go-microservice/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

// ListAll return all products from database
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

package handlers

import (
	"net/http"

	"github.com/gr8abbasi/go-microservice/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

//PUT request to update product
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := p.GetProductID(r)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err := data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not Found!", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not Found!", http.StatusInternalServerError)
		return
	}
}

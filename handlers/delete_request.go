package handlers

import (
	"net/http"

	"github.com/gr8abbasi/go-microservice/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

//DELETE product from database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := p.GetProductID(r)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not Found!", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not Found!", http.StatusInternalServerError)
		return
	}
}

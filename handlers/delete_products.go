package handlers

import (
	"net/http"

	"github.com/gr8abbasi/go-microservice/data"
)

//DELETE product from database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Product")

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

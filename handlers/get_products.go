package handlers

import (
	"net/http"

	"github.com/gr8abbasi/go-microservice/data"
)

//GET request to get all products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

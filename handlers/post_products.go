package handlers

import (
	"net/http"

	"github.com/gr8abbasi/go-microservice/data"
)

//POST request to add product
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

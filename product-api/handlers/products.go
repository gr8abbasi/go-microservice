package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gr8abbasi/go-microservice/product-api/data"
)

type Products struct {
	l *log.Logger
	v *data.Validation
}

// Returns new Products handler
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

type KeyProduct struct{}

// Return product ID from URL
func (p *Products) GetProductID(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

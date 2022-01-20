// Package classification of Product API
//
// Documentation for Product API
//
//		Schemes: http
// 		BasePath: /
// 		Version: 1.0.0
//
// 		Consumes:
// 		- application/json
//
// 		Produces:
// 		- application/json
//
// swagger:meta
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

// Returns new Products handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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

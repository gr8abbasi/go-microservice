package main

import (
	"testing"

	"github.com/gr8abbasi/go-microservice/product-api/sdk/client"
	"github.com/gr8abbasi/go-microservice/product-api/sdk/client/products"
	"github.com/stretchr/testify/assert"
)

// Test our swagger generated Client
func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8081")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	assert.GreaterOrEqual(t, len(prods.GetPayload()), 1)
}

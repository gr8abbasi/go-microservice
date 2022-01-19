package data

import "testing"

func TestValidationSuccess(t *testing.T) {
	p := &Product{
		Name:  "Test Successful",
		Price: 2.75,
		SKU:   "coffee-1",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err, "Test Failed")
	}
}

func TestValidationFail(t *testing.T) {
	p := &Product{
		Name:  "",
		Price: 0,
		SKU:   "coffee",
	}

	err := p.Validate()
	if err == nil {
		t.Fatal("Test Failed")
	}
}

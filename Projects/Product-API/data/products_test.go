package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Malhar",
		Description: "This is five letter word",
		Price: 1.23,
		SKU: "abc-bch-sdj",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

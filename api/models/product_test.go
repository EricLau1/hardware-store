package models

import "testing"

func TestProductModel_Validate(t *testing.T) {

	p := &Product{}
	p.Name = ""
	p.Price = 1600.56
	p.Quantity = 10
	p.Status = ProductStatus_Available

	expected := "product.name can't be empty"

	if err := p.Validate(); err.Error() != expected {
		t.Errorf("expected: %s, got: %s", expected, err.Error())
	}
}

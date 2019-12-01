package models

import "testing"

func TestCategoryModel_Validate(t *testing.T) {

	c := &Category{}
	c.Description = ""

	err := c.Validate()

	expected := "category.description can't empty"

	if err.Error() != expected {
		t.Errorf("expected: %s, got: %s", expected, err.Error())
	}
}

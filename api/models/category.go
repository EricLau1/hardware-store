package models

import "errors"

type Category struct {
	Model
	Description string     `gorm:"size:256;not null;unique" json:"description"`
	Products    []*Product `gorm:"foreignkey:CategoryID" json:"products"`
}

var (
	ErrCategoryEmptyDescription = errors.New("category.description can't be empty")
)

func (c *Category) Validate() error {
	if c.Description == "" {
		return ErrCategoryEmptyDescription
	}
	return nil
}

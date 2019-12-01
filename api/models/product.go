package models

import "errors"

type ProductStatus uint8

const (
	ProductStatus_Unavailable = 0
	ProductStatus_Available   = 1
)

type Product struct {
	Model
	Name       string        `gorm:"size:512;not null;unique" json:"name"`
	Price      float64       `gorm:"type:decimal(10,2);not null;default:0.0" json:"price"`
	Quantity   uint16        `gorm:"default:0;unsigned" json:"quantity"`
	Status     ProductStatus `gorm:"char(1);default:0" json:"status"`
	CategoryID uint64        `gorm:"not null" json:"category_id"`
}

var (
	ErrProductEmptyName = errors.New("product.name can't be empty")
)

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductEmptyName
	}
	return nil
}

func (p *Product) CheckStatus() {
	p.Status = ProductStatus_Unavailable
	if p.Quantity > 0 {
		p.Status = ProductStatus_Available
	}
}

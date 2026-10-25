package entity

import (
	"stockflow/internal/dto"
	"time"
)

type Product struct {
	ID             int        `db:"id"`
	Sku            string     `db:"sku"`
	Name           string     `db:"name"`
	Price          string     `db:"price"`
	Stock          int        `db:"stock"`
	IsActive       bool       `db:"is_active"`
	CreatedAt      time.Time  `db:"created_at"`
	LastMovementAt *time.Time `db:"last_movement_at"`
}

func (p *Product) ToResponseType() dto.Product {
	return dto.Product{
		ID:             p.ID,
		Sku:            p.Sku,
		Name:           p.Name,
		Price:          p.Price,
		Stock:          p.Stock,
		IsActive:       p.IsActive,
		CreatedAt:      p.CreatedAt,
		LastMovementAt: p.LastMovementAt,
	}
}

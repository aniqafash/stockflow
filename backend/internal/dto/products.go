package dto

import "time"

type Product struct {
	ID             int       `json:"id"`
	Sku            string    `json:"sku"`
	Name           string    `json:"name"`
	Price          string    `json:"price"`
	Stock          int       `json:"stock"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	LastMovementAt time.Time `json:"last_movement_at"`
}

type ProductListRequest struct {
	Page
	Search string `form:"search"`
}

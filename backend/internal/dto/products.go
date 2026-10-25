package dto

import "time"

type Product struct {
	ID             int        `json:"id"`
	Sku            string     `json:"sku"`
	Name           string     `json:"name"`
	Price          string     `json:"price"`
	Stock          int        `json:"stock"`
	IsActive       bool       `json:"isActive"`
	CreatedAt      time.Time  `json:"createdAt"`
	LastMovementAt *time.Time `json:"lastMovementAt"`
}

type ProductListRequest struct {
	Page
	Search string `form:"search"`
}

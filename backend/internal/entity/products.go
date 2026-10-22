package entity

import "time"

type Product struct {
	ID             int       `db:"id"`
	Sku            string    `db:"sku"`
	Name           string    `db:"name"`
	Price          string    `db:"price"`
	Stock          int       `db:"price"`
	IsActive       bool      `db:"is_active"`
	CreatedAt      time.Time `db:"created_at"`
	LastMovementAt time.Time `db:"last_movement_at"`
}

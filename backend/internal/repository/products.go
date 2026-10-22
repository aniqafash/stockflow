package repository

import (
	"context"
	"stockflow/internal/entity"
)

type ProductRepository struct {
	db dbtx
}

func NewProductRepository(db dbtx) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) List(ctx context.Context, search string, page entity.Page) ([]entity.Product, error) {
	var products = make([]entity.Product, 0, 999)
	const query = `
		SELECT id, sku, name, price, stock, is_active, created_at, last_movement_at
		FROM products p
		WHERE name ILIKE $1
		ORDER BY created_at
		LIMIT $2 OFFSET $3`
	var offset = (page.Page - 1) * page.Limit
	err := r.db.SelectContext(ctx, &products, query, search, page.Limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

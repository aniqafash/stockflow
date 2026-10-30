package repository

import (
	"regexp"
	"stockflow/internal/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestProductRepository_List(t *testing.T) {
	var products = []entity.Product{
		{
			ID:             1,
			Sku:            "SPT-BLUE-F40",
			Name:           "Flatshoes",
			Price:          "220000",
			Stock:          27,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: parseTime(2026, 8, 23, 0, 16, 32),
		},
		{
			ID:             2,
			Sku:            "SPT-PINK-F39",
			Name:           "Flatshoes",
			Price:          "215000",
			Stock:          18,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: parseTime(2026, 8, 15, 19, 36, 0),
		},
		{
			ID:             3,
			Sku:            "SPT-BLUE-F39",
			Name:           "Flatshoes",
			Price:          "215000",
			Stock:          12,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             4,
			Sku:            "SPT-GREN-F40",
			Name:           "Flatshoes",
			Price:          "199000",
			Stock:          76,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             5,
			Sku:            "BLS-RED-XL",
			Name:           "Blouse Lengan Balon",
			Price:          "275000",
			Stock:          11,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: parseTime(2026, 8, 22, 13, 47, 10),
		},
		{
			ID:             6,
			Sku:            "BLS-BRWN-S",
			Name:           "Blouse Lengan Balon",
			Price:          "265000",
			Stock:          0,
			IsActive:       false,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             7,
			Sku:            "BLS-RED-M",
			Name:           "Blouse Lengan Balon",
			Price:          "215000",
			Stock:          8,
			IsActive:       false,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             8,
			Sku:            "KLT-PINK-XS",
			Name:           "Kulot Plisket",
			Price:          "129000",
			Stock:          19,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             9,
			Sku:            "KLT-WHTE-XL",
			Name:           "Kulot Plisket",
			Price:          "129000",
			Stock:          213,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             10,
			Sku:            "KLT-BLUE-S",
			Name:           "Kulot Plisket",
			Price:          "129000",
			Stock:          41,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             11,
			Sku:            "KLT-GRAY-M",
			Name:           "Kulot Plisket",
			Price:          "129000",
			Stock:          68,
			IsActive:       false,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: parseTime(2026, 8, 22, 13, 47, 10),
		},
		{
			ID:             12,
			Sku:            "JKT-WHTE-2XL",
			Name:           "Jaket Bomber",
			Price:          "215000",
			Stock:          11,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: parseTime(2026, 8, 17, 17, 5, 21),
		},
		{
			ID:             13,
			Sku:            "JKT-BLCK-L",
			Name:           "Jaket Bomber",
			Price:          "319000",
			Stock:          24,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: parseTime(2026, 8, 17, 17, 5, 21),
		},
		{
			ID:             14,
			Sku:            "JKT-GRAY-S",
			Name:           "Jaket Bomber",
			Price:          "325000",
			Stock:          125,
			IsActive:       false,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             15,
			Sku:            "JKT-BLUE-XL",
			Name:           "Jaket Bomber",
			Price:          "312000",
			Stock:          27,
			IsActive:       false,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
		{
			ID:             16,
			Sku:            "JKT-BLCK-M",
			Name:           "Jaket Bomber",
			Price:          "318000",
			Stock:          133,
			IsActive:       false,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: parseTime(2026, 8, 17, 17, 5, 21),
		},
		{
			ID:             17,
			Sku:            "JKT-BLUE-L",
			Name:           "Jaket Bomber",
			Price:          "325000",
			Stock:          34,
			IsActive:       true,
			CreatedAt:      time.Date(2026, 8, 1, 0, 0, 0, 0, time.FixedZone("", 7*3600)),
			LastMovementAt: nil,
		},
	}

	tests := []struct {
		name    string
		search  string
		page    entity.Page
		want    []entity.Product
		wantErr error
	}{
		{
			name:    "list with base paging and  without search",
			search:  "",
			page:    entity.Page{Page: 1, Limit: 50},
			want:    products,
			wantErr: nil,
		},
		{
			name:    "list with paging and without search",
			search:  "",
			page:    entity.Page{Page: 2, Limit: 5},
			want:    products[5:9],
			wantErr: nil,
		},
		{
			name:    "list with search and paging",
			search:  "Jaket",
			page:    entity.Page{Page: 1, Limit: 50},
			want:    products[11:],
			wantErr: nil,
		},
		{
			name:    "list with zero results of search",
			search:  "zzz",
			page:    entity.Page{Page: 1, Limit: 50},
			want:    []entity.Product{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			db := sqlx.NewDb(sqlDB, "sqlmock")

			expectedRows := sqlmock.NewRows([]string{
				"id", "sku", "name", "price", "stock", "is_active", "created_at", "last_movement_at",
			})

			for _, row := range tt.want {
				expectedRows.AddRow(
					row.ID,
					row.Sku,
					row.Name,
					row.Price,
					row.Stock,
					row.IsActive,
					row.CreatedAt,
					row.LastMovementAt,
				)
			}

			var offset = (tt.page.Page - 1) * tt.page.Limit

			mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT id, sku, name, price, stock, is_active, created_at, last_movement_at 
				FROM products p 
				WHERE $1::text = '' OR name ILIKE $1 
				ORDER BY created_at 
				LIMIT $2 OFFSET $3`)).
				WithArgs(tt.search, tt.page.Limit, offset).
				WillReturnRows(expectedRows)

			r := NewProductRepository(db)
			got, gotErr := r.List(t.Context(), tt.search, tt.page)

			require.NoError(t, gotErr)
			require.Equal(t, got, tt.want)
		})
	}
}

func parseTime(y int, m time.Month, d, h, min, s int) *time.Time {
	t := time.Date(y, m, d, h, min, s, 0, time.FixedZone("", 7*3600))
	return &t
}

package service

import (
	"errors"
	"stockflow/internal/dto"
	"stockflow/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var dtoProducts = []dto.Product{
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

var entityProducts = []entity.Product{
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

func TestProductService_List(t *testing.T) {
	tests := []struct {
		name        string
		search      string
		pageRequest dto.Page
		returnMock  []entity.Product
		want        []dto.Product
		wantErr     error
	}{
		{
			name:        "page limit request is nil",
			search:      "",
			pageRequest: dto.Page{Page: new(1), Limit: nil},
			returnMock:  entityProducts,
			want:        dtoProducts,
			wantErr:     nil,
		},
		{
			name:        "page page request is nil",
			search:      "",
			pageRequest: dto.Page{Page: nil, Limit: new(5)},
			returnMock:  entityProducts[:4],
			want:        dtoProducts[:4],
			wantErr:     nil,
		},
		{
			name:        "database error",
			search:      "",
			pageRequest: dto.Page{Page: nil, Limit: new(5)},
			returnMock:  nil,
			want:        nil,
			wantErr:     dto.ErrLookUp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := NewMockProductRepository(ctrl)

			switch {
			case errors.Is(tt.wantErr, dto.ErrLookUp):
				mockRepo.EXPECT().List(t.Context(), tt.search, gomock.Any()).Return(tt.returnMock, errors.New("db error")).Times(1)
			default:
				mockRepo.EXPECT().List(t.Context(), tt.search, gomock.Any()).Return(tt.returnMock, nil).Times(1)
			}
			s := NewProductService(mockRepo)

			got, gotErr := s.List(t.Context(), tt.search, tt.pageRequest)

			require.Equal(t, tt.wantErr, gotErr)
			require.Equal(t, tt.want, got)
		})
	}
}

func parseTime(y int, m time.Month, d, h, min, s int) *time.Time {
	t := time.Date(y, m, d, h, min, s, 0, time.FixedZone("", 7*3600))
	return &t
}

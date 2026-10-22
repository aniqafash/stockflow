package service

import (
	"context"
	"stockflow/internal/dto"
	"stockflow/internal/entity"
)

type ProductRepository interface {
	List(ctx context.Context, search string, page entity.Page) ([]entity.Product, error)
}

type ProductService struct {
	productRepo ProductRepository
}

func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) List(ctx context.Context, search string, pageRequest dto.Page) ([]dto.Product, error) {
	var page = entity.Page{
		Limit: dto.BaseLimit,
		Page:  dto.BasePage,
	}

	if pageRequest.Limit != nil {
		if *pageRequest.Limit > dto.MaxLimit {
			page.Limit = dto.MaxLimit
		}
		page.Limit = *pageRequest.Limit
	}

	if pageRequest.Page != nil {
		page.Page = *pageRequest.Page
	}

	res, err := s.productRepo.List(ctx, search, page)
	if err != nil {
		return nil, dto.ErrLookUp
	}

	var products = make([]dto.Product, 0, len(res))
	for _, p := range res {
		products = append(products, p.ToResponseType())
	}

	return products, nil
}

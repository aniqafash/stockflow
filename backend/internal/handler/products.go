package handler

import (
	"context"
	"net/http"
	"stockflow/internal/dto"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=products.go -destination=products_service_mock.go -package=handler
type ProductService interface {
	List(ctx context.Context, search string, pageRequest dto.Page) ([]dto.Product, error)
}

type ProductHandler struct {
	productSvc ProductService
}

func NewProductHandler(productService ProductService) *ProductHandler {
	return &ProductHandler{productSvc: productService}
}

func (h *ProductHandler) List(c *gin.Context) {
	var req dto.ProductListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		responseError(c, http.StatusBadRequest, CodeBadProductSRequest, nil)
		return
	}

	products, err := h.productSvc.List(c.Request.Context(), req.Search, req.Page)
	if err != nil {
		errDetails := mapServiceError(err)
		responseError(c, errDetails.Status, errDetails.Code, nil)
		return
	}

	responseSuccess(c, http.StatusOK, products)
}

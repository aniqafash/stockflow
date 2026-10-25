package handler

import (
	"errors"
	"net/http"
	"stockflow/internal/dto"

	"github.com/gin-gonic/gin"
)

const (
	CodeBadProductSRequest = "BAD_PRODUCTS_REQUEST"
)

func responseSuccess(c *gin.Context, status int, data any) {
	c.JSON(status, dto.SuccessResponse{
		Data: data,
	})
}

func responseError(c *gin.Context, status int, code string, details *any) {
	c.JSON(status, dto.FailResponse{
		Error: dto.FailDetails{
			Code:    code,
			Details: details,
		},
	})
}

type Error struct {
	Status  int
	Code    string
	Message string
}

func mapServiceError(err error) Error {
	switch {
	case errors.Is(err, dto.ErrLookUp):
		return Error{
			Status:  http.StatusInternalServerError,
			Code:    "ERROR_LOOKUP",
			Message: err.Error(),
		}
	default:
		return Error{
			Status:  http.StatusInternalServerError,
			Code:    "ERROR_INTERNAL_SERVER",
			Message: "error internal server",
		}
	}
}

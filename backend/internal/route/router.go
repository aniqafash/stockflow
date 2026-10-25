package route

import (
	"stockflow/internal/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Product *handler.ProductHandler
}

func NewRouter(h Handlers) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Logger(),
		gin.Recovery(),
		handler.CORS(),
	)

	api := r.Group("/api/")
	{
		api.GET("/products", h.Product.List)
	}

	return r
}

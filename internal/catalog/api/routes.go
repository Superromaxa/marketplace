package api

import (
	"github.com/Superromaxa/marketplace/internal/catalog/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, brands *handlers.BrandsHandler) {
	v1 := r.Group("/api/v1")

	v1.GET("/brands", brands.Brands) //http://localhost:9001/api/v1/brands
}

package handlers

import (
	"net/http"

	"github.com/Superromaxa/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type BrandsHandler struct {
	brands *queries.BrandsHandler
}

func NewBrandsHandler(brands *queries.BrandsHandler) *BrandsHandler {
	return &BrandsHandler{brands: brands}
}

func (bh *BrandsHandler) Brands(ctx *gin.Context) {
	brands, err := bh.brands.Handle(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"brands": brands,
	})
}

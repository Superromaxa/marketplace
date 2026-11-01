package handlers

import (
	"net/http"

	"github.com/Superromaxa/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	categories *queries.CategoryHandler
}

func NewCategoriesHandler(categories *queries.CategoryHandler) *CategoriesHandler {
	return &CategoriesHandler{categories: categories}
}

func (ch *CategoriesHandler) Categories(ctx *gin.Context) {
	categories, err := ch.categories.Handle(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

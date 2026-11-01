package handlers

import (
	"net/http"

	"github.com/Superromaxa/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type CatalogItemsHandler struct {
	catalogItems *queries.CatalogItemsHandler
}

func NewCatalogItemsHandler(catalogItems *queries.CatalogItemsHandler) *CatalogItemsHandler {
	return &CatalogItemsHandler{catalogItems: catalogItems}
}

func (ci *CatalogItemsHandler) CatalogItems(ctx *gin.Context) {
	catalogItems, err := ci.catalogItems.Handle(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"catalogItems": catalogItems,
	})
}

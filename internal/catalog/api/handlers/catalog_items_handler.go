package handlers

import (
	"net/http"

	"github.com/Superromaxa/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CatalogItemsHandler struct {
	catalogItems    *queries.CatalogItemsHandler
	catalogItemByID *queries.CatalogItemByIDHandler
}

func NewCatalogItemsHandler(
	catalogItems *queries.CatalogItemsHandler,
	catalogItemByID *queries.CatalogItemByIDHandler,
) *CatalogItemsHandler {
	return &CatalogItemsHandler{catalogItems: catalogItems, catalogItemByID: catalogItemByID}
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

func (ci *CatalogItemsHandler) CatalogItemByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong id",
		})
		return
	}

	item, err := ci.catalogItemByID.Handle(ctx.Request.Context(), id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if item == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "item not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"catalogItem": item,
	})
}

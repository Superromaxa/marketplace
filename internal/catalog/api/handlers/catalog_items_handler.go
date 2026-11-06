package handlers

import (
	"net/http"

	"github.com/Superromaxa/marketplace/internal/catalog/application/commands"
	"github.com/Superromaxa/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CatalogItemsHandler struct {
	catalogItems        *queries.CatalogItemsHandler
	catalogItemByID     *queries.CatalogItemByIDHandler
	catalogItemsByTitle *queries.CatalogItemsByTitleHandler
	catalogItemsByBrand *queries.CatalogItemsByBrandHandler

	createCatalogItem *commands.CreateCatalogItemHandler
	updateCatalogItem *commands.UpdateCatalogItemHandler
	deleteCatalogItem *commands.DeleteCatalogItemHandler
}

func NewCatalogItemsHandler(
	catalogItems *queries.CatalogItemsHandler,
	catalogItemByID *queries.CatalogItemByIDHandler,
	catalogItemsByTitle *queries.CatalogItemsByTitleHandler,
	catalogItemsByBrand *queries.CatalogItemsByBrandHandler,
	createCatalogItem *commands.CreateCatalogItemHandler,
	updateCatalogItem *commands.UpdateCatalogItemHandler,
	deleteCatalogItem *commands.DeleteCatalogItemHandler,
) *CatalogItemsHandler {
	return &CatalogItemsHandler{
		catalogItems:        catalogItems,
		catalogItemByID:     catalogItemByID,
		catalogItemsByTitle: catalogItemsByTitle,
		catalogItemsByBrand: catalogItemsByBrand,

		createCatalogItem: createCatalogItem,
		updateCatalogItem: updateCatalogItem,
		deleteCatalogItem: deleteCatalogItem,
	}
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

func (ci *CatalogItemsHandler) CatalogItemsByTitle(ctx *gin.Context) {
	title := ctx.Param("title")

	items, err := ci.catalogItemsByTitle.Handle(ctx.Request.Context(), title)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": items,
	})
}

func (ci *CatalogItemsHandler) CatalogItemsByBrand(ctx *gin.Context) {
	brand := ctx.Param("brand")

	items, err := ci.catalogItemsByBrand.Handle(ctx.Request.Context(), brand)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": items,
	})
}

func (ci *CatalogItemsHandler) CreateCatalogItem(ctx *gin.Context) {
	var cmd commands.CreateCatalogItemCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := ci.createCatalogItem.Handle(ctx.Request.Context(), cmd)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (ci *CatalogItemsHandler) UpdateCatalogItem(ctx *gin.Context) {
	var cmd commands.UpdateCatalogItemCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	success, err := ci.updateCatalogItem.Handle(ctx.Request.Context(), cmd)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "item not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}

func (ci *CatalogItemsHandler) DeleteCatalogItem(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong id",
		})
		return
	}

	success, err := ci.deleteCatalogItem.Handle(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !success {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "item not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}

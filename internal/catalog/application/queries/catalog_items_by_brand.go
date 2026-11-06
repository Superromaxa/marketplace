package queries

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
)

type CatalogItemsByBrandHandler struct {
	repo repositories.CatalogItemRepository
}

func NewCatalogItemsByBrandHandler(repo repositories.CatalogItemRepository) *CatalogItemsByBrandHandler {
	return &CatalogItemsByBrandHandler{repo: repo}
}

func (ch *CatalogItemsByBrandHandler) Handle(
	ctx context.Context,
	brand string,
) ([]entities.CatalogItem, error) {
	return ch.repo.ItemsByBrand(ctx, brand)
}

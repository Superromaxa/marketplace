package queries

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
)

type CatalogItemsByTitleHandler struct {
	repo repositories.CatalogItemRepository
}

func NewCatalogItemsByTitleHandler(repo repositories.CatalogItemRepository) *CatalogItemsByTitleHandler {
	return &CatalogItemsByTitleHandler{repo: repo}
}

func (ch *CatalogItemsByTitleHandler) Handle(
	ctx context.Context,
	title string,
) ([]entities.CatalogItem, error) {
	return ch.repo.ItemsByTitle(ctx, title)
}

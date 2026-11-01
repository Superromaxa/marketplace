package queries

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
)

type CatalogItemsHandler struct {
	repo repositories.CatalogItemRepository
}

func NewCatalogItemsHandler(repo repositories.CatalogItemRepository) *CatalogItemsHandler {
	return &CatalogItemsHandler{repo: repo}
}

func (ch *CatalogItemsHandler) Handle(ctx context.Context) ([]entities.CatalogItem, error) {
	return ch.repo.Items(ctx)
}

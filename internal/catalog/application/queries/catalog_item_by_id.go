package queries

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type CatalogItemByIDHandler struct {
	repo repositories.CatalogItemRepository
}

func NewCatalogItemByIDHandler(repo repositories.CatalogItemRepository) *CatalogItemByIDHandler {
	return &CatalogItemByIDHandler{repo: repo}
}

func (ch *CatalogItemByIDHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (*entities.CatalogItem, error) {
	return ch.repo.Item(ctx, id)
}

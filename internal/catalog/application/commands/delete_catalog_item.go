package commands

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type DeleteCatalogItemHandler struct {
	repo repositories.CatalogItemRepository
}

func NewDeleteCatalogItemHandler(repo repositories.CatalogItemRepository) *DeleteCatalogItemHandler {
	return &DeleteCatalogItemHandler{repo: repo}
}

func (i *DeleteCatalogItemHandler) Handle(ctx context.Context, id uuid.UUID) (bool, error) {
	return i.repo.Delete(ctx, id)
}

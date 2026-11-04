package repositories

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/google/uuid"
)

type CatalogItemRepository interface {
	Items(ctx context.Context) ([]entities.CatalogItem, error)
	Item(ctx context.Context, id uuid.UUID) (*entities.CatalogItem, error)
	ItemsByTitle(ctx context.Context, title string) ([]entities.CatalogItem, error)
}

type BrandRepository interface {
	Brands(ctx context.Context) ([]entities.Brand, error)
}

type CategoryRepository interface {
	Categories(ctx context.Context) ([]entities.Category, error)
}

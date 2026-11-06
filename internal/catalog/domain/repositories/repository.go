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
	ItemsByBrand(ctx context.Context, title string) ([]entities.CatalogItem, error)

	Create(ctx context.Context, item entities.CatalogItem) (entities.CatalogItem, error)
	Update(ctx context.Context, item entities.CatalogItem) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
}

type BrandRepository interface {
	Brands(ctx context.Context) ([]entities.Brand, error)
}

type CategoryRepository interface {
	Categories(ctx context.Context) ([]entities.Category, error)
}

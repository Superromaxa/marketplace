package repositories

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
)

type CatalogItemRepository interface {
	Items(ctx context.Context) ([]entities.CatalogItem, error)
}

type BrandRepository interface {
	Brand(ctx context.Context) ([]entities.Brand, error)
}

type CategoryRepository interface {
	Category(ctx context.Context) ([]entities.Category, error)
}

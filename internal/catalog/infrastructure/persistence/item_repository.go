package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/google/uuid"
)

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *itemRepository {
	return &itemRepository{db: db}
}

func (i *itemRepository) Items(ctx context.Context) ([]entities.CatalogItem, error) {
	query := `
    SELECT
        ci.id,
        ci.title,
        COALESCE(ci.short_description, ''),
		COALESCE(ci.full_description, ''),
		COALESCE(ci.image_url, ''),
        b.id,
        b.title,
        c.id,
        c.title
    FROM catalog_items AS ci
    LEFT JOIN brands AS b ON b.id = ci.brand_id
    LEFT JOIN categories AS c ON c.id = ci.category_id
    ORDER BY ci.title
`
	rows, err := i.db.QueryContext(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to load items: %w", err)
	}
	defer rows.Close()

	var items []entities.CatalogItem
	for rows.Next() {
		item, err := scanCatalogItem(rows)

		if err != nil {
			return nil, fmt.Errorf("failed to scan items: %w", err)
		}

		items = append(items, item)
	}
	return items, rows.Err()
}

func scanCatalogItem(rows *sql.Rows) (entities.CatalogItem, error) {
	var item entities.CatalogItem
	var brandID *uuid.UUID
	var brandTitle *string
	var categoryID *uuid.UUID
	var categoryTitle *string

	err := rows.Scan(
		&item.ID,
		&item.Title,
		&item.ShortDescription,
		&item.FullDescription,
		&item.ImageURL,
		&item.Price,
		&brandID,
		&brandTitle,
		&categoryID,
		&categoryTitle)

	if err != nil {
		return item, fmt.Errorf("failed to scan item: %w", err)
	}

	if brandID != nil {
		item.Brand = &entities.Brand{
			BaseEntity: entities.BaseEntity{
				ID:    *brandID,
				Title: *brandTitle,
			},
		}
	}

	if categoryID != nil {
		item.Category = &entities.Category{
			BaseEntity: entities.BaseEntity{
				ID:    *categoryID,
				Title: *categoryTitle,
			},
		}
	}

	return item, nil
}

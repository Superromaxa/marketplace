package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/google/uuid"
)

const sqlCatalogItemsQuery = `
    SELECT
        ci.id,
        ci.title,
        COALESCE(ci.short_description, ''),
		COALESCE(ci.full_description, ''),
		COALESCE(ci.image_url, ''),
		ci.price,
        b.id,
        b.title,
        c.id,
        c.title
    FROM catalog_items AS ci
    LEFT JOIN brands AS b ON b.id = ci.brand_id
    LEFT JOIN categories AS c ON c.id = ci.category_id
	`

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *itemRepository {
	return &itemRepository{db: db}
}

func (i *itemRepository) Items(ctx context.Context) ([]entities.CatalogItem, error) {
	query := sqlCatalogItemsQuery + " ORDER BY ci.title"
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

		items = append(items, *item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate items: %w", err)
	}
	return items, nil
}

func (i *itemRepository) Item(ctx context.Context, id uuid.UUID) (*entities.CatalogItem, error) {
	query := sqlCatalogItemsQuery + " WHERE ci.id = $1"
	row := i.db.QueryRowContext(ctx, query, id)

	item, err := scanCatalogItem(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

func (i *itemRepository) ItemsByTitle(ctx context.Context, title string) ([]entities.CatalogItem, error) {
	query := sqlCatalogItemsQuery + " WHERE ci.title ILIKE $1 ORDER BY ci.title"
	rows, err := i.db.QueryContext(ctx, query, "%"+title+"%")

	if err != nil {
		return nil, fmt.Errorf("failed to load items by title: %w", err)
	}
	defer rows.Close()

	items := []entities.CatalogItem{}
	for rows.Next() {
		item, err := scanCatalogItem(rows)

		if err != nil {
			return nil, fmt.Errorf("failed to scan items by title: %w", err)
		}

		items = append(items, *item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate items: %w", err)
	}
	return items, nil
}

func (i *itemRepository) ItemsByBrand(ctx context.Context, brand string) ([]entities.CatalogItem, error) {
	query := sqlCatalogItemsQuery + " WHERE b.title ILIKE $1 ORDER BY ci.title"

	rows, err := i.db.QueryContext(ctx, query, "%"+brand+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to load items by brand: %w", err)
	}
	defer rows.Close()

	items := []entities.CatalogItem{}
	for rows.Next() {
		item, err := scanCatalogItem(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan items by brand: %w", err)
		}
		items = append(items, *item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate items by brand: %w", err)
	}
	return items, nil
}

func (i *itemRepository) Create(ctx context.Context, item entities.CatalogItem) (entities.CatalogItem, error) {
	if item.ID == uuid.Nil {
		item.ID = uuid.New()
	}

	var brandID, categoryID *uuid.UUID

	if item.Brand != nil {
		brandID = &item.Brand.ID
	}
	if item.Category != nil {
		categoryID = &item.Category.ID
	}

	sqlQuery := `
	INSERT INTO catalog_items (
		id,
		title,
	    short_description,
	    full_description,
	    image_url,
	    brand_id,
	    category_id,
	    price
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	_, err := i.db.ExecContext(ctx, sqlQuery,
		item.ID, item.Title, item.ShortDescription, item.FullDescription, item.ImageURL, brandID, categoryID, item.Price,
	)

	if err != nil {
		return entities.CatalogItem{}, fmt.Errorf("failed to create item: %w", err)
	}

	return item, nil
}

func (i *itemRepository) Update(ctx context.Context, item entities.CatalogItem) (bool, error) {
	var brandID, categoryID *uuid.UUID

	if item.Brand != nil {
		brandID = &item.Brand.ID
	}
	if item.Category != nil {
		categoryID = &item.Category.ID
	}

	sqlQuery := `
	UPDATE catalog_items SET
		title = $1,
	    short_description = $2,
	    full_description = $3,
	    image_url = $4,
	    brand_id = $5,
	    category_id = $6,
	    price = $7
	WHERE id = $8;
	`

	result, err := i.db.ExecContext(ctx, sqlQuery,
		item.Title, item.ShortDescription, item.FullDescription, item.ImageURL, brandID, categoryID, item.Price, item.ID)

	if err != nil {
		return false, fmt.Errorf("failed to update item[%v]: %w", item.ID, err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get affected rows: %w", err)
	}

	return n > 0, nil
}

func (i *itemRepository) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	sqlQuery := `DELETE FROM catalog_items
       WHERE id = $1;`

	result, err := i.db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete item[%v]: %w", id, err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get affected rows: %w", err)
	}

	return n > 0, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanCatalogItem(s scanner) (*entities.CatalogItem, error) {
	var item *entities.CatalogItem = &entities.CatalogItem{}
	var brandID *uuid.UUID
	var brandTitle *string
	var categoryID *uuid.UUID
	var categoryTitle *string

	err := s.Scan(
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
		return nil, fmt.Errorf("failed to scan item: %w", err)
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

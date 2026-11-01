package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Categories(ctx context.Context) ([]entities.Category, error) {
	rows, err := c.db.QueryContext(ctx,
		`SELECT id, title
		FROM categories
		ORDER BY title`)

	if err != nil {
		return nil, fmt.Errorf("failed to load categories: %w", err)
	}

	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.ID, &category.Title); err != nil {
			return nil, fmt.Errorf("failed to scan categories: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, rows.Err()
}

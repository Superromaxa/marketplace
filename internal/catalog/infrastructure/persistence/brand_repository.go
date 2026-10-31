package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
)

type brandRepository struct {
	db *sql.DB
}

func NewBrandRepository(db *sql.DB) *brandRepository {
	return &brandRepository{db: db}
}

func (b *brandRepository) Brands(ctx context.Context) ([]entities.Brand, error) {
	rows, err := b.db.QueryContext(ctx,
		`SELECT id, title
		FROM brands
		ORDER BY title
	`)

	if err != nil {
		return nil, fmt.Errorf("failed to query brands: %w", err)
	}

	defer rows.Close()

	var brands []entities.Brand
	for rows.Next() {
		var brand entities.Brand
		if err := rows.Scan(&brand.ID, &brand.Title); err != nil {
			return nil, fmt.Errorf("failed to scan brand: %w", err)
		}
		brands = append(brands, brand)
	}

	return brands, rows.Err()
}

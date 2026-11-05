package commands

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type UpdateCatalogItemCommand struct {
	ID               uuid.UUID  `json:"id"`
	Title            string     `json:"title"`
	ShortDescription string     `json:"short_description"`
	FullDescription  string     `json:"full_description"`
	ImageURL         string     `json:"image_url"`
	BrandID          *uuid.UUID `json:"brand_id"`
	CategoryID       *uuid.UUID `json:"category_id"`
	Price            float64    `json:"price"`
}

type UpdateCatalogItemHandler struct {
	repo repositories.CatalogItemRepository
}

func NewUpdateCatalogItemHandler(repo repositories.CatalogItemRepository) *UpdateCatalogItemHandler {
	return &UpdateCatalogItemHandler{repo: repo}
}

func (h *UpdateCatalogItemHandler) Handle(ctx context.Context, cmd UpdateCatalogItemCommand) (bool, error) {
	exist, err := h.repo.Item(ctx, cmd.ID)

	if err != nil {
		return false, err
	}
	if exist == nil {
		return false, nil
	}

	item := entities.CatalogItem{
		BaseEntity: entities.BaseEntity{
			ID:    cmd.ID,
			Title: cmd.Title,
		},
		ShortDescription: cmd.ShortDescription,
		FullDescription:  cmd.FullDescription,
		ImageURL:         cmd.ImageURL,
		Price:            cmd.Price,
	}

	if cmd.BrandID != nil {
		item.Brand = &entities.Brand{
			BaseEntity: entities.BaseEntity{ID: *cmd.BrandID},
		}
	}
	if cmd.CategoryID != nil {
		item.Category = &entities.Category{
			BaseEntity: entities.BaseEntity{ID: *cmd.CategoryID},
		}
	}

	return h.repo.Update(ctx, item)
}

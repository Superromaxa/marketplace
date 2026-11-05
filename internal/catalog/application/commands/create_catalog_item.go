package commands

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
	"github.com/google/uuid"
)

type CreateCatalogItemCommand struct {
	Title            string     `json:"title"`
	ShortDescription string     `json:"short_description"`
	FullDescription  string     `json:"full_description"`
	ImageURL         string     `json:"image_url"`
	BrandID          *uuid.UUID `json:"brand_id"`
	CategoryID       *uuid.UUID `json:"category_id"`
	Price            float64    `json:"price"`
}

type CreateCatalogItemHandler struct {
	repo repositories.CatalogItemRepository
}

func NewCreateCatalogItemHandler(repo repositories.CatalogItemRepository) *CreateCatalogItemHandler {
	return &CreateCatalogItemHandler{repo: repo}
}

func (ch *CreateCatalogItemHandler) Handle(ctx context.Context, cmd CreateCatalogItemCommand) (uuid.UUID, error) {
	item := entities.CatalogItem{
		BaseEntity: entities.BaseEntity{
			ID:    uuid.New(),
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

	created, err := ch.repo.Create(ctx, item)

	if err != nil {
		return uuid.Nil, err
	}

	return created.ID, nil
}

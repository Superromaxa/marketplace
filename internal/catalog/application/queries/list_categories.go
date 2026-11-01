package queries

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
)

type CategoryHandler struct {
	repo repositories.CategoryRepository
}

func NewCategoryHandler(repo repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}

func (ch *CategoryHandler) Handle(ctx context.Context) ([]entities.Category, error) {
	return ch.repo.Categories(ctx)
}

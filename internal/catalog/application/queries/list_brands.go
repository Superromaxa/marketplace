package queries

import (
	"context"

	"github.com/Superromaxa/marketplace/internal/catalog/domain/entities"
	"github.com/Superromaxa/marketplace/internal/catalog/domain/repositories"
)

type BrandsHandler struct {
	repo repositories.BrandRepository
}

func NewBrandsHandler(repo repositories.BrandRepository) *BrandsHandler {
	return &BrandsHandler{repo: repo}
}

func (bh *BrandsHandler) Handle(ctx context.Context) ([]entities.Brand, error) {
	return bh.repo.Brands(ctx)
}

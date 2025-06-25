package service

import (
	"context"

	"math/rand/v2"

	"com.int.recipies/models"
)

type DataSource interface {
	CreateRecipie(ctx context.Context, recipie models.Recipie) (int, error)
	GetRecipie(ctx context.Context, id int) (*models.Recipie, error)
	DeleteRecipie(ctx context.Context, id int) (error)
	UpdateRecipie(ctx context.Context, recipie models.Recipie) (error)
	GetRecipiesCount(ctx context.Context) (int, error)
	SearchRecipies(ctx context.Context, query string, skip int, limit int) ([]models.Recipie, error)
	ListRecipies(ctx context.Context, skip int, limit int) ([]models.Recipie, error)
}

type RecipieService struct {
	d DataSource
}

func NewService(d DataSource) *RecipieService {
	return &RecipieService{
		d: d,
	}
}

func (s *RecipieService) GetRecipie(ctx context.Context, id int) (*models.Recipie, error) {
	recipie, err := s.d.GetRecipie(ctx, id)
	if err != nil {
		return nil, err
	}
	return recipie, nil
}

func (s *RecipieService) CreateRecipie(ctx context.Context, recipie models.Recipie) (int, error) {
	randomNumber := rand.Int()
	recipie.ID = randomNumber
	id, err := s.d.CreateRecipie(ctx, recipie)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *RecipieService) DeleteRecipie(ctx context.Context, id int) (error) {
	err := s.d.DeleteRecipie(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipieService) UpdateRecipie(ctx context.Context, recipie models.Recipie, id int) (error) {
	recipie.ID = id
	err := s.d.UpdateRecipie(ctx, recipie)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipieService) GetRecipiesPagination(ctx context.Context, skip int, limit int) (*models.RecipiePagination, error) {
	recipies, err := s.d.ListRecipies(ctx,skip, limit)
	if err != nil {
		return nil, err
	}
	count, err := s.d.GetRecipiesCount(ctx)
	if err != nil {
		return nil, err
	}
	var recipesPagination models.RecipiePagination
	recipiePaginationMeta := models.RecipiePaginationMeta{
		Total: count,
		Skipped: skip,
		Limit: limit,
	}
	recipesPagination.Meta = recipiePaginationMeta
	recipesPagination.Recipies = recipies
	return &recipesPagination, nil
}

func (s *RecipieService) SearchRecipies(ctx context.Context, query string) (*models.RecipiePagination, error) {
	skip := 0
	limit := 100
	recipies, err := s.d.SearchRecipies(ctx, query, skip, limit)
	if err != nil {
		return nil, err
	}
	count, err := s.d.GetRecipiesCount(ctx)
	if err != nil {
		return nil, err
	}
	var recipesPagination models.RecipiePagination
	recipiePaginationMeta := models.RecipiePaginationMeta{
		Total: count,
		Skipped: skip,
		Limit: limit,
	}
	recipesPagination.Meta = recipiePaginationMeta
	recipesPagination.Recipies = recipies
	return &recipesPagination, nil
}
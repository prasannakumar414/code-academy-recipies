package service

import (
	"context"

	"math/rand/v2"

	"com.int.recipies/models"
	"go.uber.org/zap"
)

type DataSource interface {
	CreateRecipie(ctx context.Context, recipie models.Recipie) (error)
	GetRecipie(ctx context.Context, id int) (*models.Recipie, error)
	DeleteRecipie(ctx context.Context, id int) (error)
	UpdateRecipie(ctx context.Context, recipie models.Recipie) (error)
	GetRecipiesCount(ctx context.Context) (int, error)
	SearchRecipies(ctx context.Context, query string, skip int, limit int) ([]models.Recipie, error)
	ListRecipies(ctx context.Context, skip int, limit int) ([]models.Recipie, error)
}

type RecipieService struct {
	d DataSource
	logger *zap.Logger
}

func NewService(d DataSource, logger *zap.Logger) *RecipieService {
	return &RecipieService{
		d: d,
		logger: logger,
	}
}

func (s *RecipieService) GetRecipie(ctx context.Context, id int) (*models.Recipie, error) {
	recipie, err := s.d.GetRecipie(ctx, id)
	if err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}
	return recipie, nil
}

func (s *RecipieService) CreateRecipie(ctx context.Context, recipie models.Recipie) (int, error) {
	randomNumber := rand.IntN(25000)
	recipie.ID = randomNumber
	err := s.d.CreateRecipie(ctx, recipie)
	if err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return -1, err
	}
	return recipie.ID, nil
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

func (s *RecipieService) SearchRecipiesPagination(ctx context.Context, query string, skip int, limit int) (*models.RecipiePagination, error) {
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
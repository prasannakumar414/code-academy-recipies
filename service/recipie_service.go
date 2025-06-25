package service

import (
	"context"

	"math/rand/v2"

	"com.int.recipies/models"
)

type DataSource interface {
	CreateRecipie(ctx context.Context, recipie models.Recipie) (int, error)
	GetRecipie(ctx context.Context, id int) (*models.Recipie, error)
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

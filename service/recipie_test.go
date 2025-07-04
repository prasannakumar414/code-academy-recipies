package service_test

import (
	"context"
	"testing"

	"com.int.recipies/models"
	"com.int.recipies/service"
	"go.uber.org/zap"
)

type db struct {
}

func (db *db) CreateRecipie(ctx context.Context, recipie models.Recipie) error {
	return nil
}
func (db *db) GetRecipie(ctx context.Context, id int) (*models.Recipie, error) {
	return nil, nil
}
func (db *db) DeleteRecipie(ctx context.Context, id int) error {
	return nil
}
func (db *db) UpdateRecipie(ctx context.Context, recipie models.Recipie) error {
	return nil
}
func (db *db) GetRecipiesCount(ctx context.Context) (int, error) {
	return 0, nil
}
func (db *db) SearchRecipies(ctx context.Context, query string, skip int, limit int) ([]models.Recipie, error) {
	return []models.Recipie{}, nil
}
func (db *db) ListRecipies(ctx context.Context, skip int, limit int) ([]models.Recipie, error) {
	return []models.Recipie{}, nil
}

func TestCreate(t *testing.T) {
	logger, _ := zap.NewProduction()
	service := service.NewService(&db{}, logger)
	id, err := service.CreateRecipie(context.Background(), models.Recipie{})
	if id < 0 {
		t.Errorf("id cannot be negative")
	}
	if err != nil {
		t.Errorf("Error when creating recipie %v", err)
	}
}

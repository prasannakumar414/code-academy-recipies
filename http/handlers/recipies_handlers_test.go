package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"com.int.recipies/http/handlers"
	"com.int.recipies/models"
	"github.com/gin-gonic/gin"
)

type mock struct {
}

func (m *mock) CreateRecipie(ctx context.Context, recipie models.Recipie) (int, error) {
	return 0, nil
}
func (m *mock) GetRecipie(ctx context.Context, id int) (*models.Recipie, error) {
	return nil, nil
}
func (m *mock) DeleteRecipie(ctx context.Context, id int) error {
	return nil
}
func (m *mock) UpdateRecipie(ctx context.Context, recipie models.Recipie, id int) error {
	return nil
}
func (m *mock) GetRecipiesPagination(ctx context.Context, skip int, limit int) (*models.RecipiePagination, error) {
	return nil, nil
}
func (m *mock) SearchRecipies(ctx context.Context, query string) (*models.RecipiePagination, error) {
	return nil, nil
}
func (m *mock) SearchRecipiesPagination(ctx context.Context, query string, skip int, limit int) (*models.RecipiePagination, error) {
	return nil, nil
}

func TestGetRecipie(t *testing.T) {
	w := httptest.NewRecorder()
	handlers := handlers.NewRecipieHandler(&mock{})
	router := gin.Default()
	router.GET("/", handlers.GetRecipie)
	req, _ := http.NewRequest("GET", "/?id=7331", nil)
	router.ServeHTTP(w, req)
	if w.Result().StatusCode != 200 {
		t.Errorf("error")
	}
}

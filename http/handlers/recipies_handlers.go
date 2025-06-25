package handlers

import (
	"context"
	"net/http"
	"strconv"

	"com.int.recipies/models"
	"github.com/gin-gonic/gin"
)

type RecipiesService interface {
	CreateRecipie(ctx context.Context, recipie models.Recipie) (int, error)
	GetRecipie(ctx context.Context, id int) (*models.Recipie, error)
	DeleteRecipie(ctx context.Context, id int) error
	UpdateRecipie(ctx context.Context, recipie models.Recipie, id int) error
	GetRecipiesPagination(ctx context.Context, skip int, limit int) (*models.RecipiePagination, error)
	SearchRecipies(ctx context.Context, query string) (*models.RecipiePagination, error)
}

type RecipiesHandlers struct {
	service RecipiesService
}

func NewRecipieHandler(RecipiesService RecipiesService) *RecipiesHandlers {
	return &RecipiesHandlers{
		service: RecipiesService,
	}
}

func (r *RecipiesHandlers) GetRecipie(c *gin.Context) {
	id := c.Query("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid recipie id"})
	}
	recipie, err := r.service.GetRecipie(c.Request.Context(), num)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipie not found"})
	}
	c.IndentedJSON(http.StatusOK, recipie)
}

func (r *RecipiesHandlers) CreateRecipie(c *gin.Context) {
	var recipie models.Recipie
	if err := c.BindJSON(&recipie); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid recipie"})
		return
	}
	id, err := r.service.CreateRecipie(c.Request.Context(), recipie)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error creating recipie"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func (r *RecipiesHandlers) DeleteRecipie(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid recipie id"})
	}
	err = r.service.DeleteRecipie(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipie not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

func (r *RecipiesHandlers) UpdateRecipie(c *gin.Context) {
	var recipie models.Recipie
	if err := c.BindJSON(&recipie); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid recipie"})
		return
	}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid recipie id"})
	}
	err = r.service.UpdateRecipie(c.Request.Context(), recipie, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipie not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

func (r *RecipiesHandlers) RecipiePagination(c *gin.Context) {
	var recipiePaginationRequest models.RecipiePaginationRequest
    limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid limit"})
	}
	recipiePaginationRequest.Limit = limit
	skip, err := strconv.Atoi(c.Query("skip"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid skip"})
	}
	recipiePaginationRequest.Skip = skip
	recipiePaginationRequest.Query = c.Query("query")

	recipiePagination, err := r.service.GetRecipiesPagination(c.Request.Context(), recipiePaginationRequest.Skip, recipiePaginationRequest.Limit)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error getting pagination"})
		return
	}
	c.IndentedJSON(http.StatusOK, recipiePagination)
}
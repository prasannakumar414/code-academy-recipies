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

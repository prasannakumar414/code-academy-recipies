package handlers

import (
	"net/http"

	"com.int.recipies/models"
	"github.com/gin-gonic/gin"
)

type RecipiesService interface {

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
	recipie := models.Recipie{
		ID: 1,
		Name: "name",
		Instructions: "ingredients",
		PrepTime: "1 hour",
		Description: "this is a recipie",
	}
	c.IndentedJSON(http.StatusOK, recipie)
}
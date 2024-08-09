package flat

import (
	"avito/internal/models"
	"avito/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FlatHandler struct {
	// TODO: add repository
	repo *repository.FlatRepository
}

func NewFlatHandler(repo *repository.FlatRepository) *FlatHandler {
	return &FlatHandler{
		repo: repo,
	}
}

func (h *FlatHandler) Create(c *gin.Context) {
	flat := models.Flat{}
	err := c.BindJSON(&flat)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read request body"})
		c.Abort()
		return
	}

	err = validateFlatData(flat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	flat, err = h.repo.Create(flat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, flat)
}

func (h *FlatHandler) Update(c *gin.Context) {

}

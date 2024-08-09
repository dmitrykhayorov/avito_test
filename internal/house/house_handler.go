package house

import (
	"avito/internal/models"
	"avito/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HouseHandler struct {
	repo *repository.HouseRepository
}

func NewHouseHandler(repo *repository.HouseRepository) *HouseHandler {
	return &HouseHandler{repo: repo}
}

func (h *HouseHandler) Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *HouseHandler) Create(c *gin.Context) {
	house := models.House{}
	err := c.BindJSON(&house)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	fmt.Println(house)
	newH, err := h.repo.Create(&house)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, newH)
}

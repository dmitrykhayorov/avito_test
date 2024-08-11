package flat

import (
	"avito/internal/models"
	"avito/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type FlatHandler struct {
	repo repository.FlatRepositoryInterface
}

func NewFlatHandler(repo repository.FlatRepositoryInterface) *FlatHandler {
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
		slog.ErrorContext(c, err.Error())
		response := models.Response500{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, flat)
}

func (h *FlatHandler) Update(c *gin.Context) {
	body := models.FlatUpdateRequestBody{}
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	err = validateUpdateRequestBody(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	currentFlatStatus, err := h.repo.GetFlatStatus(body.FlatId)

	if err != nil {
		slog.ErrorContext(c, err.Error())
		response := fmt.Sprintf("flat: %v in house %v not found", body.FlatId, body.HouseId)
		c.JSON(http.StatusBadRequest, gin.H{"message": response})
		c.Abort()
		return
	}

	if currentFlatStatus == models.StatusOnModeration {
		c.JSON(http.StatusBadRequest, gin.H{"message": "already on moderation"})
		c.Abort()
		return
	}

	updatedFlat, err := h.repo.Update(body.FlatId, body.HouseId, body.Status)

	if err != nil {
		slog.ErrorContext(c, err.Error())
		response := models.Response500{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, updatedFlat)
}

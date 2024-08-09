package house

import (
	"avito/internal/models"
	"avito/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HouseHandler struct {
	repo *repository.HouseRepository
}

func NewHouseHandler(repo *repository.HouseRepository) *HouseHandler {
	return &HouseHandler{repo: repo}
}

func (h *HouseHandler) GetFlatsByHouseID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// TODO: add proper check
	userRole, _ := c.Get("userRole")
	u, ok := userRole.(string)
	if !ok {
		return
	}
	flats, err := h.repo.GetFlatsByHouseID(models.UserRole(u), uint32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, flats)
}

func (h *HouseHandler) Create(c *gin.Context) {
	house := models.House{}
	err := c.BindJSON(&house)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err = validateHouseData(house)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	house, err = h.repo.Create(house)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, house)
}

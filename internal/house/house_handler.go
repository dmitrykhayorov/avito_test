package house

import (
	"avito/internal/models"
	"avito/internal/repository"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type HouseHandler struct {
	repo repository.HouseRepositoryInterface
}

func NewHouseHandler(repo repository.HouseRepositoryInterface) *HouseHandler {
	return &HouseHandler{repo: repo}
}

func (h *HouseHandler) GetFlatsByHouseID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	u, _ := c.Get("userRole")
	userRole, ok := u.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot get user role"})
		c.Abort()
		return
	}

	flats, err := h.repo.GetFlatsByHouseID(models.UserRole(userRole), id)

	if err != nil {
		slog.ErrorContext(c, err.Error())
		response := models.Response500{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
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
		slog.ErrorContext(c, err.Error())
		response := models.Response500{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, house)
}

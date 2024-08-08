package house

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HouseHandler struct {
}

func NewHouseHandler() *HouseHandler {
	return &HouseHandler{}
}

func (h *HouseHandler) Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *HouseHandler) Create(c *gin.Context) {

}

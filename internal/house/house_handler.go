package house

import "github.com/gin-gonic/gin"

type HouseHandler struct {
}

func NewHouseHandler() *HouseHandler {
	return &HouseHandler{}
}

func (h *HouseHandler) Get(c *gin.Context) {

}

func (h *HouseHandler) Create(c *gin.Context) {

}

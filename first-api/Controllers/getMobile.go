package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMobiles ... Get all mobile devices
// @Summary Get All Mobile Devices
// @Description Retrieve all mobile devices
// @ID get-all-mobiles
// @Accept json
// @Produce json
// @Success 200 {array} models.Mobile
// @Failure 404 {object} gin.H
// @Router /mobiles [get]
func (h *Handler) GetMobiles(c *gin.Context) {
	mobile, err := h.Service.GetAllMobiles()
	// mobile, err := models.GetAllMobiles()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mobile)
	}
}

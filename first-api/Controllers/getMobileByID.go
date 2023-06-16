package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMobileByID ... Get a mobile device by ID
// @Summary Get Mobile Device by ID
// @Description Retrieve a mobile device by ID
// @ID get-mobile-by-id
// @Accept json
// @Produce json
// @Param id path int true "Mobile device ID"
// @Success 200 {object} models.Mobile
// @Failure 404 {object} gin.H
// @Router /mobile/{id} [get]
func (h *Handler) GetMobileByID(c *gin.Context) {
	id := c.Params.ByName("id")
	mobile, err := h.Service.GetMobileByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mobile)
	}
}

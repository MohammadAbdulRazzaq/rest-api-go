package Controllers

import (
	"first-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateMobile ... Update the mobile device
// @Summary Update Mobile Device
// @Description Update a mobile device by ID
// @ID update-mobile
// @Accept json
// @Produce json
// @Param id path int true "Mobile device ID"
// @Param mobile body models.Mobile true "Mobile device object"
// @Success 200 {object} models.Mobile
// @Failure 404 {object} gin.H
// @Router /mobile/{id} [put]
func (h *Handler) UpdateMobile(c *gin.Context) {
	var user models.Mobile
	id := c.Params.ByName("id")
	mobile, err := h.Service.GetMobileByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&mobile)
	updatedmobile, err := h.Service.UpdateMobile(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, updatedmobile)
	}
}

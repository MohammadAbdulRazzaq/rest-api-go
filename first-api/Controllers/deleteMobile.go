package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteMobile ... Delete the mobile device
// @Summary Delete Mobile Device
// @Description Delete a mobile device by ID
// @ID delete-mobile
// @Accept json
// @Produce json
// @Param id path int true "Mobile device ID"
// @Success 200 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /mobile/{id} [delete]
func (h *Handler) DeleteMobile(c *gin.Context) {
	id := c.Params.ByName("id")
	success, err := h.Service.DeleteMobile(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else if success {
		c.JSON(http.StatusOK, gin.H{"id": id, "message": "is deleted"})
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

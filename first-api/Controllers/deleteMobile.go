package Controllers

import (
	"first-api/Models"
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
func DeleteMobile(c *gin.Context) {
	var user Models.Mobile
	id := c.Params.ByName("id")
	err := Models.DeleteMobile(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

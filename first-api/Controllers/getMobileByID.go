package Controllers

import (
	"first-api/Models"
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
// @Success 200 {object} Models.Mobile
// @Failure 404 {object} gin.H
// @Router /mobile/{id} [get]
func GetMobileByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Mobile
	err := Models.GetMobileByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

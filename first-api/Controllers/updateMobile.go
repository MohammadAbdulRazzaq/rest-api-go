package Controllers

import (
	"first-api/Models"
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
// @Param mobile body Models.Mobile true "Mobile device object"
// @Success 200 {object} Models.Mobile
// @Failure 404 {object} gin.H
// @Router /mobile/{id} [put]
func UpdateMobile(c *gin.Context) {
	var user Models.Mobile
	id := c.Params.ByName("id")
	err := Models.GetMobileByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateMobile(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

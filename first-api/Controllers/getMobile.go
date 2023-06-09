package Controllers

import (
	"first-api/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMobiles ... Get all mobile devices
// @Summary Get All Mobile Devices
// @Description Retrieve all mobile devices
// @ID get-all-mobiles
// @Accept json
// @Produce json
// @Success 200 {array} Models.Mobile
// @Failure 404 {object} gin.H
// @Router /mobiles [get]
func GetMobiles(c *gin.Context) {
	var user []Models.Mobile
	err := Models.GetAllMobiles(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

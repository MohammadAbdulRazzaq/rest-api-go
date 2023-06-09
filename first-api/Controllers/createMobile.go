package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMobile ... Create User
// @Summary Create User
// @Description Create a new mobile device
// @ID create-mobile
// @Accept json
// @Produce json
// @Param mobile body Mobile true "Mobile object to be created"
// @Success 200 {object} Mobile
// @Failure 404 {object} map[string]interface{}
// @Router /mobile [post]
func CreateMobile(c *gin.Context) {
	var user Models.Mobile
	c.BindJSON(&user)
	err := Models.CreateMobile(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

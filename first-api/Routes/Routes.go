package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("/mobile-api")
	{
		grp1.GET("mobile", Controllers.GetMobiles)
		grp1.POST("mobile", Controllers.CreateMobile)
		grp1.GET("mobile/:id", Controllers.GetMobileByID)
		grp1.PUT("mobile/:id", Controllers.UpdateMobile)
		grp1.DELETE("mobile/:id", Controllers.DeleteMobile)
	}
	return r
}

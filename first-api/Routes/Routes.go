package Routes

import (
	"first-api/Controllers"
	"first-api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("/api")
	{
		grp1.POST("/signin", Controllers.SignIn)
	}
	protected := r.Group("/admin")
	{
		protected.Use(middleware.JwtAuthMiddleware())
		protected.GET("mobile", Controllers.GetMobiles)
		protected.POST("mobile", Controllers.CreateMobile)
		protected.GET("mobile/:id", Controllers.GetMobileByID)
		protected.PUT("mobile/:id", Controllers.UpdateMobile)
		protected.DELETE("mobile/:id", Controllers.DeleteMobile)
	}

	return r
}

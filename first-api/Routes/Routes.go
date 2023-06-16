package Routes

import (
	"first-api/Controllers"
	"first-api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	publicgrp := r.Group("/api")
	{
		publicgrp.POST("/signin", Controllers.SignIn)
	}
	protected := r.Group("/admin")
	{
		h := Controllers.New()
		protected.Use(middleware.JwtAuthMiddleware())
		protected.GET("mobile", h.GetMobiles)
		protected.POST("mobile", h.CreateMobile)
		protected.GET("mobile/:id", h.GetMobileByID)
		protected.PUT("mobile/:id", h.UpdateMobile)
		protected.DELETE("mobile/:id", h.DeleteMobile)
	}

	return r
}

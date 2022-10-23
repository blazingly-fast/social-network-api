package routes

import (
	"github.com/blazingly-fast/social-network-api/controllers"
	"github.com/blazingly-fast/social-network-api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.Use(middleware.Authenticate())
	r.GET("/users", controllers.GetUsers())
	r.GET("/users/:user_id", controllers.GetUser())
}

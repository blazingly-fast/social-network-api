package main

import (
	"github.com/blazingly-fast/social-network-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

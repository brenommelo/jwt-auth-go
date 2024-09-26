package main

import (
	"net/http"

	"github.com/brenommelo/jwt-auth-go/controllers"
	"github.com/brenommelo/jwt-auth-go/initializers"
	"github.com/brenommelo/jwt-auth-go/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

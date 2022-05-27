package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(app *Application, router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// ping: always return 200 OK
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		// register endpoint
		v1.POST("/register", app.RegisterController.Register)
		// authentication endpoints
		auth := v1.Group("/auth")
		{
			auth.POST("/login", app.AuthMiddleware.LoginHandler)
			auth.POST("/refresh_token", app.AuthMiddleware.RefreshHandler)
		}
		// user endpoints
		user := v1.Group("/user")
		user.Use(app.AuthMiddleware.MiddlewareFunc())
		{
			user.POST("", app.UserController.Create)
			user.GET("/:id", app.UserController.GetByID)
			user.GET("", app.UserController.GetAll)
			user.PUT("/:id", app.UserController.Update)
			user.DELETE("/:id", app.UserController.Delete)
		}
		// about topic endpoints
		aboutTopic := v1.Group("/about_topic")
		aboutTopic.Use(app.AuthMiddleware.MiddlewareFunc())
		{
			aboutTopic.GET("/:id", app.AboutTopicController.GetByID)
			aboutTopic.GET("", app.AboutTopicController.GetAll)
		}

	}

	router.NoRoute()
}

package infrastructure

import (
	"github.com/ITK13201/portfolio/backend/ent"
	"github.com/ITK13201/portfolio/backend/interfaces/controllers"
	"github.com/ITK13201/portfolio/backend/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Application struct {
	AuthMiddleware     middlewares.AuthMiddleware
	UserController     controllers.UserController
	RegisterController controllers.RegisterController
}

func BuildApplication(sqliClient *ent.Client) *Application {
	authMiddleware := middlewares.NewAuthMiddleware(sqliClient)
	userController := controllers.NewUserController(sqliClient)
	registerController := controllers.NewRegisterController(sqliClient)

	return &Application{
		AuthMiddleware:     authMiddleware,
		UserController:     userController,
		RegisterController: registerController,
	}
}

func SetUp(sqlClient *ent.Client) *gin.Engine {
	engine := gin.Default()
	engine.Use(corsHandler())
	app := BuildApplication(sqlClient)
	InitRouter(app, engine)

	return engine
}

func corsHandler() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

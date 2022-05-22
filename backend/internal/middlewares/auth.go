package middlewares

import (
	"github.com/ITK13201/portfolio/backend/config"
	"github.com/ITK13201/portfolio/backend/ent"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

const (
	Realm               = "private zone"
	IdentityKeyUsername = "username"
	IdentityKeyID       = "id"
)

type AuthMiddleware interface {
	LoginHandler(c *gin.Context)
	MiddlewareFunc() gin.HandlerFunc
	RefreshHandler(c *gin.Context)
}

type authMiddleware struct {
	ginJWTMiddleware *jwt.GinJWTMiddleware
}

func NewAuthMiddleware(sqlClient *ent.Client) AuthMiddleware {
	cfg := *config.Cfg
	authMiddlewareHandler := authMiddlewareHandler{sqlClient: sqlClient}

	ginJWTMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           Realm,
		Key:             []byte(cfg.SecretKey),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     IdentityKeyUsername,
		PayloadFunc:     authMiddlewareHandler.payloadFunc,
		IdentityHandler: authMiddlewareHandler.identityHandler,
		Authenticator:   authMiddlewareHandler.authenticator,
		Authorizator:    authMiddlewareHandler.authorizator,
		Unauthorized:    authMiddlewareHandler.unauthorized,
		LoginResponse:   authMiddlewareHandler.loginResponse,
		RefreshResponse: authMiddlewareHandler.refreshResponse,
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return &authMiddleware{
		ginJWTMiddleware: ginJWTMiddleware,
	}
}

func (a *authMiddleware) LoginHandler(c *gin.Context) {
	a.ginJWTMiddleware.LoginHandler(c)
}

func (a *authMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return a.ginJWTMiddleware.MiddlewareFunc()
}

func (a *authMiddleware) RefreshHandler(c *gin.Context) {
	a.ginJWTMiddleware.RefreshHandler(c)
}

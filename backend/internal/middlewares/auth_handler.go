package middlewares

import (
	"context"
	"github.com/ITK13201/portfolio/backend/domain"
	"github.com/ITK13201/portfolio/backend/ent"
	"github.com/ITK13201/portfolio/backend/ent/user"
	crypto "github.com/ITK13201/portfolio/backend/internal/crypt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type authUser struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type authMiddlewareHandler struct {
	sqlClient     *ent.Client
	authUser      *authUser
	retrievedUser *ent.User
}

// Checks if the user is valid in the db
func (a *authMiddlewareHandler) authenticator(c *gin.Context) (interface{}, error) {
	var loginVals domain.Login
	var err error
	if err = c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	ctx := context.Background()
	a.retrievedUser, err = a.sqlClient.User.Query().Where(user.UsernameEQ(strings.ToLower(loginVals.Username))).Only(ctx)

	// User do not exist
	if ent.IsNotFound(err) {
		return "", jwt.ErrFailedAuthentication
	}

	// Has valid password?
	if crypto.IsHashedPasswordEqualWithPlainPassword(a.retrievedUser.Password, loginVals.Password) {
		a.authUser = &authUser{
			ID:       a.retrievedUser.ID,
			Username: a.retrievedUser.Username,
		}
		return a.authUser, nil
	}

	return "", jwt.ErrFailedAuthentication
}

// Returns the payload for the token if the user is valid
func (a *authMiddlewareHandler) payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*authUser); ok {
		return jwt.MapClaims{
			IdentityKeyID:       v.ID,
			IdentityKeyUsername: v.Username,
		}
	}
	return jwt.MapClaims{}
}

// Once the token is received in a request and is valid, it extracts the data
func (a *authMiddlewareHandler) identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &authUser{
		ID:       int64(claims[IdentityKeyID].(float64)),
		Username: claims[IdentityKeyUsername].(string),
	}
}

// At this moment, the token is valid
// Saves user info in context
func (a *authMiddlewareHandler) authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*authUser); ok {
		c.Set(IdentityKeyID, v.ID)
		c.Set(IdentityKeyUsername, v.Username)
	}

	return true
}

// Message to show when the provided token is not valid
func (a *authMiddlewareHandler) unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func (a *authMiddlewareHandler) refreshResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, &domain.RefreshResponse{
		Token:      token,
		Expiration: expire.Format(time.RFC3339),
	})
}

func (a *authMiddlewareHandler) loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, &domain.LoginResponse{
		Token:      token,
		Expiration: expire.Format(time.RFC3339),
		User:       *a.retrievedUser,
	})
}

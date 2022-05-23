package controllers

import (
	"github.com/ITK13201/portfolio/backend/domain"
	"github.com/ITK13201/portfolio/backend/ent"
	"github.com/ITK13201/portfolio/backend/interfaces/interactors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController interface {
	Register(c *gin.Context)
}

type registerController struct {
	sqlClient          *ent.Client
	registerInteractor interactors.RegisterInteractor
}

func NewRegisterController(sqlClient *ent.Client) RegisterController {
	return &registerController{
		sqlClient:          sqlClient,
		registerInteractor: interactors.NewRegisterInteractor(sqlClient),
	}
}

func (r *registerController) Register(c *gin.Context) {
	var registerVals domain.Register
	err := c.Bind(&registerVals)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	u, err := r.registerInteractor.CreateUser(registerVals)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

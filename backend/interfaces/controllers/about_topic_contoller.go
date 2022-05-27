package controllers

import (
	"github.com/ITK13201/portfolio/backend/ent"
	"github.com/ITK13201/portfolio/backend/interfaces/interactors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AboutTopicController interface {
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type aboutTopicController struct {
	sqlClient            *ent.Client
	aboutTopicInteractor interactors.AboutInteractor
}

func NewAboutTopicController(sqlClient *ent.Client) AboutTopicController {
	return &aboutTopicController{
		sqlClient:            sqlClient,
		aboutTopicInteractor: interactors.NewAboutInteractor(sqlClient),
	}
}

func (controller *aboutTopicController) GetByID(c *gin.Context) {
	id64, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	topic, err := controller.aboutTopicInteractor.GetAboutTopicByID(id64)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, topic)
}

func (controller *aboutTopicController) GetAll(c *gin.Context) {
	topics, err := controller.aboutTopicInteractor.GetAllAboutTopics()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, topics)
}

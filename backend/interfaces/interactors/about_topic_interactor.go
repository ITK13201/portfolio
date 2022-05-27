package interactors

import (
	"context"
	"github.com/ITK13201/portfolio/backend/domain"
	"github.com/ITK13201/portfolio/backend/ent"
	"github.com/ITK13201/portfolio/backend/ent/abouttopic"
)

type AboutInteractor interface {
	GetAboutTopicByID(id int64) (*domain.AboutTopic, error)
	GetAllAboutTopics() ([]*domain.AboutTopic, error)
}

type aboutInteractor struct {
	sqlClient *ent.Client
}

func NewAboutInteractor(sqlClient *ent.Client) AboutInteractor {
	return &aboutInteractor{
		sqlClient: sqlClient,
	}
}

func parseAboutTopic(topic *ent.AboutTopic) *domain.AboutTopic {
	return &domain.AboutTopic{
		Title:         topic.Title,
		DescriptionJp: topic.DescriptionJp,
		DescriptionEn: topic.DescriptionEn,
		ImagePath:     topic.Edges.Image.Path,
		Priority:      topic.Priority,
	}
}

func (controller *aboutInteractor) GetAboutTopicByID(id int64) (*domain.AboutTopic, error) {
	ctx := context.Background()
	topic, err := controller.sqlClient.AboutTopic.Query().Where(abouttopic.IDEQ(id)).WithImage().Only(ctx)
	if err != nil {
		return nil, err
	}
	parsedTopic := parseAboutTopic(topic)
	return parsedTopic, nil
}

func (controller *aboutInteractor) GetAllAboutTopics() ([]*domain.AboutTopic, error) {
	ctx := context.Background()
	topics, err := controller.sqlClient.AboutTopic.Query().WithImage().All(ctx)
	var parsedTopics []*domain.AboutTopic
	for i := 0; i < len(topics); i++ {
		parsedTopics = append(parsedTopics, parseAboutTopic(topics[i]))
	}
	if err != nil {
		return nil, err
	}
	return parsedTopics, nil
}

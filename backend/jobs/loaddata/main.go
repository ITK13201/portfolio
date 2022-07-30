package loaddata

import (
	"context"
	"github.com/ITK13201/portfolio/backend/ent"
	"github.com/ITK13201/portfolio/backend/infrastructure"
	"github.com/ITK13201/portfolio/backend/usecase"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

type LoaddataFlags struct {
	File  string
	Type  string
	Model string
}

var Flags *LoaddataFlags

func Run(cmd *cobra.Command, args []string) {
	flags := *Flags
	filePath := flags.File
	formatType := flags.Type
	modelName := flags.Model

	var data []map[string]string
	var err error
	switch formatType {
	case "csv":
		data, err = usecase.LoadCsv(filePath)
		if err != nil {
			log.Fatal(err)
		}
	case "json":
		data, err = usecase.LoadJson(filePath)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unexpected format type.")
	}

	sqlClient := infrastructure.NewSqlClient()
	ctx := context.Background()

	switch modelName {
	case "works":
		storeWorks(data, sqlClient, ctx)
	case "about_topics":
		storeAboutTopics(data, sqlClient, ctx)
	case "images":
		storeImages(data, sqlClient, ctx)
	}
}

func storeWorks(data []map[string]string, sqlClient *ent.Client, ctx context.Context) {
	bulk := make([]*ent.WorkCreate, len(data))
	for i, entity := range data {
		id, err := strconv.ParseInt(entity["id"], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		title := entity["title"]
		descJp := entity["description_jp"]
		descEn := entity["description_en"]
		image_id, err := strconv.ParseInt(entity["image_id"], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		link := entity["link"]
		priority, err := strconv.Atoi(entity["priority"])
		if err != nil {
			log.Fatal(err)
		}
		bulk[i] = sqlClient.Work.Create().SetID(id).SetTitle(title).SetDescriptionJp(descJp).SetDescriptionEn(descEn).SetImageID(image_id).SetLink(link).SetPriority(priority)
	}
	_, err := sqlClient.Work.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func storeAboutTopics(data []map[string]string, sqlClient *ent.Client, ctx context.Context) {
	bulk := make([]*ent.AboutTopicCreate, len(data))
	for i, entity := range data {
		id, err := strconv.ParseInt(entity["id"], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		title := entity["title"]
		descJp := entity["description_jp"]
		descEn := entity["description_en"]
		image_id, err := strconv.ParseInt(entity["image_id"], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		priority, err := strconv.Atoi(entity["priority"])
		if err != nil {
			log.Fatal(err)
		}
		bulk[i] = sqlClient.AboutTopic.Create().SetID(id).SetTitle(title).SetDescriptionJp(descJp).SetDescriptionEn(descEn).SetImageID(image_id).SetPriority(priority)
	}
	_, err := sqlClient.AboutTopic.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func storeImages(data []map[string]string, sqlClient *ent.Client, ctx context.Context) {
	bulk := make([]*ent.ImageCreate, len(data))
	for i, entity := range data {
		id, err := strconv.ParseInt(entity["id"], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		slug := entity["slug"]
		path := entity["path"]
		bulk[i] = sqlClient.Image.Create().SetID(id).SetSlug(slug).SetPath(path)
	}
	_, err := sqlClient.Image.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

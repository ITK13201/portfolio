package infrastructure

import (
	"log"

	"github.com/ITK13201/portfolio/backend/config"
	"github.com/ITK13201/portfolio/backend/ent"
	_ "github.com/go-sql-driver/mysql"
)

func NewSqlClient() *ent.Client {
	cfg := *config.Cfg

	entOptions := []ent.Option{}
	if cfg.Debug {
		entOptions = append(entOptions, ent.Debug())
	}

	client, err := ent.Open("mysql", cfg.Dsn, entOptions...)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

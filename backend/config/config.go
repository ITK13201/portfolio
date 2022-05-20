package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/xo/dburl"
)

type Config struct {
	Debug       bool   `json:"debug" default:"false"`
	DatabaseUrl string `json:"database-url" required:"true"`
	DatabaseDsn string `json:"database-dsn" ignored:"true"`
	Port        int    `json:"port" required:"true"`
}

var Cfg *Config
var JST *time.Location
var UTC *time.Location

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	db, err := dburl.Parse(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	cfg.DatabaseDsn = db.DSN

	return &cfg, err
}

func init() {
	var err error

	Cfg, err = LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	JST, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	UTC, err = time.LoadLocation("UTC")
	if err != nil {
		log.Fatal(err)
	}
}

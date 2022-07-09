package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/xo/dburl"
)

const Version string = "1.0.0"

type Config struct {
	// config from environment variables
	Debug       bool   `json:"debug" default:"false" envconfig:"DEBUG"`
	DatabaseUrl string `json:"database-url" required:"true" envconfig:"DATABASE_URL"`
	Port        string `json:"port" required:"true" envconfig:"PORT"`
	SecretKey   string `json:"secret-key" required:"true" envconfig:"SECRET_KEY"`
	// not environment variables
	Dsn string `json:"dsn" ignored:"true"`
}

var Cfg *Config
var JST *time.Location
var UTC *time.Location

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}

func AddDsn(cfg *Config) (*Config, error) {
	db, err := dburl.Parse(cfg.DatabaseUrl)
	cfg.Dsn = db.DSN
	return cfg, err
}

func init() {
	var err error
	var cfg *Config

	cfg, err = LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = AddDsn(cfg)
	if err != nil {
		log.Fatal(err)
	}
	Cfg = cfg

	JST, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	UTC, err = time.LoadLocation("UTC")
	if err != nil {
		log.Fatal(err)
	}
}
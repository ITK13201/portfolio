package cmd

import (
	"fmt"

	"github.com/ITK13201/portfolio/backend/config"
	"github.com/ITK13201/portfolio/backend/infrastructure"
)

func Run() {
	cfg := *config.Cfg

	router := infrastructure.InitRouter()

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}

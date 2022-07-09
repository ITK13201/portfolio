package main

import (
	"github.com/ITK13201/portfolio/backend/cmd"
	_ "github.com/ITK13201/portfolio/backend/cmd/job"
)

func main() {
	cmd.Execute()
}

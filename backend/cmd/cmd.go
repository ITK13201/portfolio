package cmd

import (
	"context"
	"fmt"
	"github.com/ITK13201/portfolio/backend/infrastructure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ITK13201/portfolio/backend/config"
)

func Run() {
	cfg := *config.Cfg
	sqlClient := infrastructure.NewSqlClient()

	engine := infrastructure.SetUp(sqlClient)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

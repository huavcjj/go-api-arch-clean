package main

import (
	"context"
	"errors"
	"fmt"
	"go-api-arch-clean/infrastructure/database"
	"go-api-arch-clean/infrastructure/web"
	"go-api-arch-clean/pkg"
	"go-api-arch-clean/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	appEnv := pkg.GetEnvDefault("APP_ENV", "development")
	if appEnv == "development" {
		err := godotenv.Load(".env.development")
		if err != nil {
			logger.Warn("Error loading .env.local file")
		}
	}

	db, err := database.NewDatabaseSQLFactory(database.InstanceMySQL)
	if err != nil {
		logger.Fatal(err.Error())
	}

	server, err := web.NewServer(web.InstanceGin, db)
	if err != nil {
		logger.Fatal(err.Error())
	}
	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}
	<-ctx.Done()
}

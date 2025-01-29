package main

import (
	"context"
	"errors"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/cromerc/projectBase/internal/v1/adapter/logger"
	"github.com/cromerc/projectBase/internal/v1/port"
	"github.com/cromerc/projectBase/internal/v1/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	// If the file does not exist, continue since the environment variables might still be set
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	var cfg service.Config
	opts := env.Options{
		RequiredIfNoDef: true,
	}

	cfg, err = env.ParseAsWithOptions[service.Config](opts)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	log := logger.New(cfg.Log)
	log.Ctx(ctx).Info("Logger initialized")

	server := port.NewServer(log)
	hs := service.NewHTTPServer(context.Background(), log, server)
	err = hs.Run()
	if err != nil {
		log.Fatal(err)
	}
}

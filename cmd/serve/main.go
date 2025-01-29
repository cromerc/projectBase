package main

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/cromerc/projectBase/internal/v1/adapter/logger"
	"github.com/cromerc/projectBase/internal/v1/port"
	"github.com/cromerc/projectBase/internal/v1/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	var cfg service.Config
	opts := env.Options{
		RequiredIfNoDef: true,
	}
	err = env.ParseWithOptions(&cfg, opts)
	if err != nil {
		panic(err)
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

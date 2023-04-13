package main

import (
	"context"

	"github.com/caarlos0/env/v6"
	"github.com/khusainnov/lamoda/internal/app"
	"github.com/khusainnov/lamoda/internal/config"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewProduction()

	cfg := &config.Config{
		L: l,
	}

	if err := env.Parse(cfg); err != nil {
		cfg.L.Fatal("cannot parse config", zap.Error(err))
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		cfg.L.Fatal("error due run the server", zap.Error(err))
	}

}

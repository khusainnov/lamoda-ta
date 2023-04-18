package app

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/khusainnov/lamoda/internal/app/productservice"
	"github.com/khusainnov/lamoda/internal/config"
	"github.com/khusainnov/lamoda/internal/db"
	"github.com/khusainnov/lamoda/internal/repository"
	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *config.Config) error {
	lis, err := net.Listen("tcp", cfg.RPCAddr)
	if err != nil {
		return fmt.Errorf("cannot create listener, %w", err)
	}

	dbClient, err := db.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("error due connect to db, %w", err)
	}

	repo := repository.NewRepository(dbClient.GetDB())
	productService := productservice.NewService(cfg.L, repo)

	if err = rpc.Register(productService); err != nil {
		return fmt.Errorf("cannot register product service, %w", err)
	}

	cfg.L.Info("start listening rpc", zap.String("PORT", cfg.RPCAddr))
	for {
		conn, err := lis.Accept()
		if err != nil {
			return fmt.Errorf("cannot accept connections")
		}

		go func(conn net.Conn) {
			defer conn.Close()
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

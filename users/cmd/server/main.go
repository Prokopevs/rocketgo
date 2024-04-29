package main

import (
	"context"
	"fmt"
	"net"
	"os"
	// "sync"

	"github.com/Prokopevs/rocketgo/schema"
	"github.com/Prokopevs/rocketgo/users/internal/core"
	"github.com/Prokopevs/rocketgo/users/internal/pg"
	"github.com/Prokopevs/rocketgo/users/internal/server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// PG_CONN=hello GRPC_SERVER_ADDR=hello2 ./server

const (
	exitCodeInitError = 2
)
func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	d, err := pg.Connect(context.Background(), cfg.pgConnString)
	if err != nil {
		return err
	}

	service := core.NewServiceImpl(d)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	// wg := &sync.WaitGroup{}
	listener, err := net.Listen("tcp", cfg.grpcAddr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	grpcService := server.NewGRPC(service)
	schema.RegisterUsersServer(grpcServer, grpcService)

	sugaredLogger.Infow("GRPC server is starting.", "addr", listener.Addr())
	
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitCodeInitError)
	}
}

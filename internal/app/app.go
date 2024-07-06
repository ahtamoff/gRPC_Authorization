package app

import (
	grpcapp "gRPC_Authorization/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	gRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	//TODO: инициализировать хранилище

	//TODO: init auth service
	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		gRPCServer: grpcApp,
	}
}

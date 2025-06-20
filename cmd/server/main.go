// инициализация DB и запуск gRPC
package main

import (
	"github.com/Saneckeg/users-service/internal/database"
	"github.com/Saneckeg/users-service/internal/transport/grpc"
	"github.com/Saneckeg/users-service/internal/user"
	"log"
)

func main() {
	database.InitDB()
	repo := user.NewUserRepository(database.DB)
	svc := user.NewService(repo)

	if err := grpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}

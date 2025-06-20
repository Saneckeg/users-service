// настройка и запуск grpc.Server
package grpc

import (
	userpb "github.com/Saneckeg/project-protos/proto/user"
	"github.com/Saneckeg/users-service/internal/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RunGRPC(svc *user.UserService) error {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	log.Println("gRPC server listening on :50051")

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	reflection.Register(grpcServer)

	return nil
}

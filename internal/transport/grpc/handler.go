// реализация всех gRPC-методов
package grpc

import (
	"context"
	userpb "github.com/Saneckeg/project-protos/proto"
	"github.com/Saneckeg/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

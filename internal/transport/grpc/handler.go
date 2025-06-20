// реализация всех gRPC-методов
package grpc

import (
	"context"

	userpb "github.com/Saneckeg/project-protos/proto/user"
	"github.com/Saneckeg/users-service/internal/user"
)

type Handler struct {
	svc *user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u := user.User{
		Email: req.Email,
	}

	createdUser, err := h.svc.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    createdUser.Id,
			Email: createdUser.Email,
		},
	}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {

	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userspb := make([]*userpb.User, len(users))

	for _, user := range users {

		userspb = append(userspb, &userpb.User{
			Id:    user.Id,
			Email: user.Email,
		})
	}

	return &userpb.ListUsersResponse{Users: userspb}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {

	updatedUser, err := h.svc.UpdateUserByID(uint(req.User.Id), user.User{
		Email: req.User.Email,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    updatedUser.Id,
			Email: updatedUser.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {

	_, err := h.svc.DeleteUserByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &userpb.DeleteUserResponse{}, nil

}

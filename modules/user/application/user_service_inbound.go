package application

import (
	"context"

	"example/modules/user/port/inbound"
	"example/modules/user/port/inbound/feature"
)

type UserService struct {
	createUser feature.CreateUserUseCase
}

func NewUserService(createUser feature.CreateUserUseCase) inbound.UserServiceInbound {
	return &UserService{
		createUser: createUser,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req feature.CreateUserServiceInboundInput) (*feature.CreateUserServiceInboundOutput, error) {
	//...

	return nil, nil
}

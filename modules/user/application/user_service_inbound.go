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

func (s *UserService) CreateUser(ctx context.Context, inputS feature.CreateUserServiceInboundInput) (*feature.CreateUserServiceInboundOutput, error) {
	inputUC := feature.CreateUserUseCaseInput{
		Name:     inputS.Name,
		Email:    inputS.Email,
		Password: inputS.Password,
	}

	outputUC, err := s.createUser.Execute(ctx, inputUC)
	if err != nil {
		return nil, err
	}

	outputS := &feature.CreateUserServiceInboundOutput{
		ID:        outputUC.User.ID.String(),
		Name:      outputUC.User.Name,
		Email:     outputUC.User.Email.String(),
		CreatedAt: outputUC.User.CreatedAt.String(),
		UpdatedAt: outputUC.User.UpdatedAt.String(),
	}

	//TODO: Dispatch UserCreated event

	return outputS, nil
}

package inbound

import "example/modules/user/port/inbound/features"

// Inbound / Service

type UserServiceInbound interface {
	CreateUserServiceInbound
}

// Inbound / Service / Create User
type CreateUserServiceInbound interface {
	features.CreateUserService
}

type CreateUserServiceInboundInput struct {
	features.CreateUserServiceRequest
}

type CreateUserServiceInboundOutput struct {
	features.CreateUserServicePresenter
}

// Inbound / Repository

type UserRepositoryInbound interface {
	features.CreateUserRepository
}

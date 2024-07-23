package inbound

import "example/modules/user/port/inbound/feature"

// Inbound / Service

type UserServiceInbound interface {
	CreateUserServiceInbound
}

// Inbound / Service / Create User
type CreateUserServiceInbound interface {
	feature.CreateUserService
}

type CreateUserServiceInboundInput struct {
	feature.CreateUserServiceRequest
}

type CreateUserServiceInboundOutput struct {
	feature.CreateUserServicePresenter
}

// Inbound / Repository

type UserRepositoryInbound interface {
	feature.CreateUserRepository
}

package outbound

import "example/modules/user/port/inbound"

// Outbound / Service

type UserServiceOutbound interface {
	inbound.CreateUserServiceInbound
}

// Outbound / Service / Create User

type CreateUserServiceOutboundInput struct {
	inbound.CreateUserServiceInboundInput
}

type CreateUserServiceOutboundOutput struct {
	inbound.CreateUserServiceInboundOutput
}

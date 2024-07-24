package inbound

import (
	"example/modules/user/port/inbound/feature"
)

// Inbound / Service
type UserServiceInbound interface {
	feature.CreateUserServiceInbound
}

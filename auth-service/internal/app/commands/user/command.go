package commands_user

import (
	v1 "github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

type UpdateProfileCommand struct {
	FirstName string
	LastName  string
	Phone     string
	Address   string
}

type UpdateProfileCommandResponse struct {
	*v1.StatusMessage
}

type GetProfileCommand struct{}

type GetProfileCommandResponse struct {
	FirstName string
	LastName  string
	Phone     string
	Address   string
	Email     string
	*v1.StatusMessage
}

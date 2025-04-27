package userdetaildto

import "github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"

type CreateUserDetailCommand struct {
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
}

type CreateUserDetailResponse struct {
	*statusmsg.StatusMessage
}

type UpdateUserDetailCommand struct {
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
}

type UpdateUserDetailResponse struct {
	*statusmsg.StatusMessage
}

type GetUserDetailCommand struct {
	UserGuid string
}

type GetUserDetailResponse struct {
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
	*statusmsg.StatusMessage
}

type DeleteUserDetailCommand struct {
	UserGuid string
}

type DeleteUserDetailResponse struct {
	*statusmsg.StatusMessage
}

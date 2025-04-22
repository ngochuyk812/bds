package commands_auth

import (
	v1 "github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
)

type LoginCommand struct {
	Email    string
	Password string
}
type LoginCommandResponse struct {
	AccessToken  string
	RefreshToken string
	*v1.StatusMessage
}

type SignUpCommand struct {
	LastName  string
	FirstName string
	Email     string
	Password  string
}
type SignUpCommandResponse struct {
	*v1.StatusMessage
}

type VerifySignUpCommand struct {
	Email string
	Otp   string
}
type VerifySignUpCommandResponse struct {
	*v1.StatusMessage
}

type RefreshTokenCommand struct {
	RefreshToken string
}
type RefreshTokenCommandResponse struct {
	AccessToken  string
	RefreshToken string
	*v1.StatusMessage
}

type LogoutCommand struct{}

type LogoutCommandResponse struct {
	*v1.StatusMessage
}

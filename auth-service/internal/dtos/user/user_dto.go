package userdto

import "github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"

type UpdateProfileCommand struct {
	FirstName string
	LastName  string
	Phone     string
	Address   string
}

type UpdateProfileResponse struct {
	*statusmsg.StatusMessage
}

type GetProfileCommand struct{}

type GetProfileResponse struct {
	FirstName string
	LastName  string
	Phone     string
	Address   string
	Email     string
	*statusmsg.StatusMessage
}

type LoginCommand struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	*statusmsg.StatusMessage
}

type SignUpCommand struct {
	LastName  string
	FirstName string
	Email     string
	Password  string
}

type SignUpResponse struct {
	*statusmsg.StatusMessage
}

type VerifySignUpCommand struct {
	Email string
	Otp   string
}

type VerifySignUpResponse struct {
	*statusmsg.StatusMessage
}

type RefreshTokenCommand struct {
	RefreshToken string
}

type RefreshTokenResponse struct {
	AccessToken  string
	RefreshToken string
	*statusmsg.StatusMessage
}

type LogoutCommand struct{}

type LogoutResponse struct {
	*statusmsg.StatusMessage
}

package connectrpc

import (
	"auth_service/internal/infra"

	"github.com/ngochuyk812/proto-bds/gen/auth/v1/authv1connect"
)

var _ authv1connect.AuthServiceHandler = &authServerHandler{}

type authServerHandler struct {
	cabin infra.Cabin
}

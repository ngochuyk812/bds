package connectrpc

import (
	"auth_service/internal/config"
	"auth_service/internal/infra"
	"auth_service/internal/infra/global"
	"auth_service/internal/usecase"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ngochuyk812/building_block/interceptors"
	"github.com/ngochuyk812/proto-bds/gen/auth/v1/authv1connect"
)

var _ authv1connect.AuthServiceHandler = &authServerHandler{}

type authServerHandler struct {
	usecaseManager usecase.UsecaseManager
	cabin          infra.Cabin
}

func NewAuthServer(usecaseManager usecase.UsecaseManager, cabin infra.Cabin) (pattern string, handler http.Handler) {
	impl := &authServerHandler{
		usecaseManager: usecaseManager,
		cabin:          cabin,
	}
	path, handler := authv1connect.NewAuthServiceHandler(impl,
		connect.WithInterceptors(
			interceptors.NewAuthInterceptor(config.SecretKey, &global.PoliciesPath),
			interceptors.NewLoggingInterceptor(),
		),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
	)
	return path, handler
}

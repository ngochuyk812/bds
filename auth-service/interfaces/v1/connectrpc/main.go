package connectrpc

import (
	"auth_service/internal/infra"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ngochuyk812/building_block/interceptors"
	"github.com/ngochuyk812/proto-bds/gen/auth/v1/authv1connect"
)

var _ authv1connect.AuthServiceHandler = &authServerHandler{}

type authServerHandler struct {
	cabin infra.Cabin
}

func NewAuthServer(cabin infra.Cabin) (pattern string, handler http.Handler) {
	impl := &authServerHandler{
		cabin: cabin,
	}
	path, handler := authv1connect.NewAuthServiceHandler(impl,
		connect.WithInterceptors(
			interceptors.NewAuthInterceptor(cabin.GetInfra().GetConfig().SecretKey, cabin.GetInfra().GetConfig().PoliciesPath),
			interceptors.NewLoggingInterceptor(cabin.GetInfra().GetLogger()),
		),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
	)
	return path, handler
}

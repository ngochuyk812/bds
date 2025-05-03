package connectrpc

import (
	"net/http"
	"property_service/internal/config"
	"property_service/internal/infra"
	"property_service/internal/infra/global"

	"connectrpc.com/connect"
	"github.com/ngochuyk812/building_block/interceptors"
	"github.com/ngochuyk812/proto-bds/gen/property/v1/propertyv1connect"
)

var _ propertyv1connect.PropertyServiceHandler = &propertyServerHandler{}

type propertyServerHandler struct {
	cabin infra.Cabin
}

func NewPropertyServer(cabin infra.Cabin) (pattern string, handler http.Handler) {
	impl := &propertyServerHandler{
		cabin: cabin,
	}
	path, handler := propertyv1connect.NewPropertyServiceHandler(impl,
		connect.WithInterceptors(
			interceptors.NewAuthInterceptor(config.SecretKey, &global.PoliciesPath),
			interceptors.NewLoggingInterceptor(),
		),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
	)
	return path, handler
}

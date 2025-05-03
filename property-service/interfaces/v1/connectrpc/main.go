package connectrpc

import (
	"net/http"
	"property_service/internal/infra"

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
			interceptors.NewAuthInterceptor(cabin.GetInfra().GetConfig().SecretKey, cabin.GetInfra().GetConfig().PoliciesPath),
			interceptors.NewLoggingInterceptor(cabin.GetInfra().GetLogger()),
		),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
	)
	return path, handler
}

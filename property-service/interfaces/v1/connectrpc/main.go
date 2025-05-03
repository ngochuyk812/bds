package connectrpc

import (
	"property_service/internal/infra"
	"property_service/internal/usecases"

	"github.com/ngochuyk812/proto-bds/gen/property/v1/propertyv1connect"
)

var _ propertyv1connect.PropertyServiceHandler = &propertyServerHandler{}

type propertyServerHandler struct {
	cabin    infra.Cabin
	useCases usecases.UsecaseManager
}

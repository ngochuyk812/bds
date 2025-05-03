package connectrpc

import (
	"context"

	"connectrpc.com/connect"
	propertyv1 "github.com/ngochuyk812/proto-bds/gen/property/v1"
)

// CreateProperty implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) CreateProperty(ctx context.Context, req *connect.Request[propertyv1.CreatePropertyRequest]) (*connect.Response[propertyv1.CreatePropertyResponse], error) {
	panic("unimplemented")
}

// DeleteProperty implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) DeleteProperty(context.Context, *connect.Request[propertyv1.DeletePropertyRequest]) (*connect.Response[propertyv1.DeletePropertyResponse], error) {
	panic("unimplemented")
}

// FetchProperties implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) FetchProperties(context.Context, *connect.Request[propertyv1.FetchPropertiesRequest]) (*connect.Response[propertyv1.FetchPropertiesResponse], error) {
	panic("unimplemented")
}

// UpdateProperty implements propertyv1connect.PropertyServiceHandler.
func (p *propertyServerHandler) UpdateProperty(context.Context, *connect.Request[propertyv1.UpdatePropertyRequest]) (*connect.Response[propertyv1.UpdatePropertyResponse], error) {
	panic("unimplemented")
}

package interceptorsinternal

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"

	"connectrpc.com/connect"
	"github.com/ngochuyk812/proto-bds/gen/statusmsg/v1"
	utilsv1 "github.com/ngochuyk812/proto-bds/gen/utils/v1"
)

func RecovertInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (resp connect.AnyResponse, err error) {
			resRecover := &utilsv1.BaseResponse{
				Status: &statusmsg.StatusMessage{},
			}
			defer func() {
				if r := recover(); r != nil {
					log.Printf("panic recovered in ConnectRPC: %v\n%s", r, debug.Stack())
					resRecover.Status.Code = statusmsg.StatusCode_STATUS_CODE_INTERNAL_ERROR
					err = errFromPanic(r)
					if err != nil {
						resRecover.Status.Extras = []string{err.Error()}
					}
					resp = connect.NewResponse[utilsv1.BaseResponse](resRecover)
				}
			}()

			resp, err = next(ctx, req)
			return resp, err
		}
	}
}

func errFromPanic(r any) error {
	switch x := r.(type) {
	case string:
		return fmt.Errorf("panic: %s", x)
	case error:
		return fmt.Errorf("panic: %w", x)
	default:
		return fmt.Errorf("panic: %v", x)
	}
}

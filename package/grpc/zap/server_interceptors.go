package grpc_zap

import (
	"context"
	"path"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// SystemField is used in every client-side log statement.
	SystemField = zap.String("system", "grpc")

	// ServerField is used in every server-side log statement.
	ServerField = zap.String("span.kind", "server")
)

type Validator interface {
	Validate() error
}

// UnaryServerInterceptor returns a new unary server interceptors
func UnaryServerInterceptor(logger *zap.Logger, opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateServerOpt(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqID := WithRequestID(ctx)

		ctx = requestid.NewContext(ctx, reqID)
		startTime := time.Now()

		newCtx := newLoggerForCall(ctx, logger, info.FullMethod, startTime, reqID)
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		resp, err := handler(newCtx, req)
		if !o.shouldLog(info.FullMethod, err) {
			return resp, err
		}

		return resp, err
	}
}

func serverCallFields(fullMethodString string) []zapcore.Field {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)
	return []zapcore.Field{
		SystemField,
		ServerField,
		zap.String("grpc.service", service),
		zap.String("grpc.method", method),
	}
}

func newLoggerForCall(
	ctx context.Context,
	logger *zap.Logger,
	fullMethodString string,
	start time.Time,
	reqID string,
) context.Context {
	var f []zapcore.Field
	f = append(f, zap.String("grpc.start_time", start.Format(time.RFC3339)))
	f = append(f, zap.String("request_id", reqID))
	if d, ok := ctx.Deadline(); ok {
		f = append(f, zap.String("grpc.request.deadline", d.Format(time.RFC3339)))
	}
	callLog := logger.With(append(f, serverCallFields(fullMethodString)...)...)
	return ctxzap.ToContext(ctx, callLog)
}

package grpc_zap

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log := ctxzap.Extract(ctx)

		ctx = metadata.AppendToOutgoingContext(ctx, "request_id", WithRequestID(ctx))

		log.Sugar().Infow("starting grpc request", "method", method)

		defer func() {
			log.Info("grpc request finished", []zapcore.Field{
				zap.String("method", method),
				zap.Float32("latency", float32(time.Since(time.Now()).Nanoseconds()/1000)/1000),
			}...)
		}()

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

package grpc_zap

import (
	"context"

	"github.com/google/uuid"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
)

const (
	DefaultRequestIDKey  = "X-Request-ID"
	TransferRequestIDKey = "request_id"
)

// WithRequestID either extracts a existing and valid request ID from the context or generates a new one
func WithRequestID(ctx context.Context) (reqID string) {
	reqID, exists := FromContext(ctx)
	if !exists || reqID == "" {
		reqID := newRequestID()
		return reqID
	}
	return reqID
}

func newRequestID() string {
	return uuid.New().String()
}

func FromContext(ctx context.Context) (string, bool) {
	if reqID, ok := gateway.Header(ctx, DefaultRequestIDKey); ok {
		return reqID, ok
	}
	if reqID, ok := gateway.Header(ctx, TransferRequestIDKey); ok {
		return reqID, ok
	}

	return "", false
}

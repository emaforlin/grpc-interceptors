package interceptors

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryLoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		elapsed := time.Since(start)

		st, _ := status.FromError(err)
		logger.Info("Handling request",
			zap.String("service", info.FullMethod),
			zap.Duration("duration", elapsed),
			zap.String("status", st.Code().String()),
		)

		return resp, err
	}
}

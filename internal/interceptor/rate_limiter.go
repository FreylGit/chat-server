package interceptor

import (
	"chat-server/internal/rate_limiter"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type RateLimiterInterceptor struct {
	rateLimiter *rate_limiter.TokenBucketLimiter
}

func NewRateLimiterInterceptor(limiter *rate_limiter.TokenBucketLimiter) *RateLimiterInterceptor {
	return &RateLimiterInterceptor{rateLimiter: limiter}
}

func (r *RateLimiterInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !r.rateLimiter.Allow() {
		return nil, errors.New("too many request")
	}

	return handler(ctx, req)
}

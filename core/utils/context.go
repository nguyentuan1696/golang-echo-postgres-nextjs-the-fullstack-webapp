package utils

import (
    "context"
    "time"
)

const defaultTimeout = 10 * time.Second

// WithTimeout wraps a context with a timeout and returns the new context and cancel function
// If timeout is <= 0, it uses defaultTimeout
// If parent context is nil, it uses context.Background()
func WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
    if ctx == nil {
        ctx = context.Background()
    }
    
    if timeout <= 0 {
        timeout = defaultTimeout
    }
    
    return context.WithTimeout(ctx, timeout)
}

// WithDeadline is similar to WithTimeout but accepts a specific deadline time
func WithDeadline(ctx context.Context, deadline time.Time) (context.Context, context.CancelFunc) {
    if ctx == nil {
        ctx = context.Background()
    }
    
    return context.WithDeadline(ctx, deadline)
}
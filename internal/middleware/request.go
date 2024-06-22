package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asdavies/auth/internal/assert"
	"github.com/google/uuid"
)

// Key to use when setting the request ID in the context
type key int

const RequestIDKey key = 0

// RequestIDMiddleware is a middleware that adds a request ID to the context of each request
func RequestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := uuid.New().String()
        ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
        r = r.WithContext(ctx)
        w.Header().Set("X-Request-ID", requestID)
        next.ServeHTTP(w, r)
    })
}

// FromContext retrieves the request ID from the context
func FromContext(ctx context.Context) string {
    requestID, ok := ctx.Value(RequestIDKey).(string)
    assert.PanicIf(!ok, fmt.Sprintf("expected a string value for RequestIDKey in context, but got %T. Ensure the context is initialized correctly.", ctx.Value(RequestIDKey)))
    return requestID
}


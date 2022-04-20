package middlewares

// https://github.com/ascarter/requestid
import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int

const RIDKey ctxKey = ctxKey(0)

// NewContext creates a context with request id.
func NewContext(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, RIDKey, rid)
}

// FromContext returns the request id from context.
func FromContext(ctx context.Context) (string, bool) {
	rid, ok := ctx.Value(RIDKey).(string)
	return rid, ok
}

// RequestIDHandler sets unique request id.
// If header `X-Request-ID` is already present in the request, that is considered the
// request id. Otherwise, generates a new unique ID.
func RequestIDHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get("X-Request-ID")
		if rid == "" {
			rid = uuid.New().String()
			r.Header.Set("X-Request-ID", rid)
		}
		ctx := NewContext(r.Context(), rid)
		h(w, r.WithContext(ctx))
	}
}

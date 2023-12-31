package middleware

import (
	"context"
	"net/http"

	"github.com/blreynolds4/go-common/examples/web/internal/sys/app"
	"github.com/blreynolds4/go-common/log"
	"github.com/blreynolds4/go-common/web"
)

// Auth simulates a midware for authentication.
func Auth(next web.Handler) web.Handler {

	// Wrap this handler around the next one provided.
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(app.KeyValues).(*app.Values)

		log.Dev(v.TraceID, "Auth", "******> Authorized")
		v.Auth = "1234"

		return next(ctx, w, r, params)
	}
}

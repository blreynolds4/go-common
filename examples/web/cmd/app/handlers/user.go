// Package handlers contains the handler logic for processing requests.
package handlers

import (
	"context"
	"net/http"

	"github.com/blreynolds4/go-common/examples/web/internal/sys/app"
	"github.com/blreynolds4/go-common/examples/web/internal/user"
	"github.com/blreynolds4/go-common/web"
)

// UserList returns all the existing users in the system.
// 200 Success, 404 Not Found, 500 Internal
func UserList(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	u, err := user.List(ctx, v.TraceID, v.DB)
	if err != nil {
		web.Error(ctx, w, err)
		return err
	}

	return web.Respond(ctx, w, u, http.StatusOK)
}

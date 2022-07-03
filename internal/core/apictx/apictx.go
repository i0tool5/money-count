package apictx

import (
	"context"
)

// UserCtx represents user context in request
type UserCtx string

// User returns user context
func User(ctx context.Context) interface{} {
	var ct UserCtx
	return ctx.Value(ct)
}

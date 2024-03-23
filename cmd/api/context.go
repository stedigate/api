package api

import (
	"context"
	"log/slog"
	"net/http"
)

type contextKey string

const requestIdContextKey = contextKey("requestId")
const slogFields = contextKey("slogFields")

type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying
// handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	return h.Handler.Handle(ctx, r)
}

// AppendCtx adds an slog attribute to the provided context so that it will be
// included in any Record created with such context
func AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	if v, ok := parent.Value(slogFields).([]slog.Attr); ok {
		v = append(v, attr)
		return context.WithValue(parent, slogFields, v)
	}

	v := []slog.Attr{}
	v = append(v, attr)
	return context.WithValue(parent, slogFields, v)
}

// const userContextKey = contextKey("user")
// func (app *application) contextSetUser(r *http.Request, user *models.User) *http.Request {
// 	ctx := context.WithValue(r.Context(), userContextKey, user)
// 	return r.WithContext(ctx)
// }
//
// func (app *application) contextGetUser(r *http.Request) *models.User {
// 	user, ok := r.Context().Value(userContextKey).(*models.User)
// 	if !ok {
// 		panic("missing user value in request context")
// 	}
//
// 	return user
// }

func (app *application) contextSetRequestId(r *http.Request, requestId string) *http.Request {
	ctx := context.WithValue(r.Context(), requestIdContextKey, requestId)
	ctx = AppendCtx(ctx, slog.String("RequestID", requestId))
	return r.WithContext(ctx)
}

func (app *application) contextGetRequestId(r *http.Request) string {
	requestId, ok := r.Context().Value(requestIdContextKey).(string)
	if !ok {
		panic("missing requestId value in request context")
	}

	return requestId
}

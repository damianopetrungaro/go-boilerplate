package trace

import "context"

type (
	id struct{}
)

//Value Retrieve the trace id given a context
func Value(ctx context.Context) interface{} {
	return ctx.Value(id{})
}

//WithValue Give back a new context with trace id
func WithValue(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, id{}, traceID)
}

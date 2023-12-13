package accesslog

import "context"

type resolverT struct{}

var (
	resolverKey = resolverT{}
)

// newResolverContext - creates a new Context with a resolver
func newResolverContext(ctx context.Context, fn resolverFunc) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if fn == nil {
		return ctx
	}
	return context.WithValue(ctx, resolverKey, fn)
}

// resolveFromContext - return a resolved resource from a Context
func resolveFromContext(ctx context.Context, rsc string) string {
	if ctx == nil {
		return rsc
	}
	i := ctx.Value(resolverKey)
	if i == nil {
		return rsc
	}
	if fn, ok := i.(resolverFunc); ok {
		return fn(rsc)
	}
	return rsc
}

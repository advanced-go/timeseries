package access1

import (
	"context"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

type queryFunc[T pgxsql.Scanner[T]] func(context.Context, http.Header, string, string, map[string][]string, ...any) ([]T, *core.Status)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, query queryFunc[T]) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	if query == nil {
		query = pgxsql.QueryT[T] //(ctx, h, accessLogResource, accessLogSelect, values)
	}
	//entries, status = pgxsql.QueryT[Entry](ctx, h, accessLogResource, accessLogSelect, values)
	entries, status = query(ctx, h, accessLogResource, accessLogSelect, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, h2, status
	}
	//t, status = pgxsql.Scan[Entry](rows)
	//if !status.OK() {
	//	e.Handle(status, core.RequestId(h))
	//	return nil, h2, status
	//}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

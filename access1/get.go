package access1

import (
	"context"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	entries, status = pgxsql.QueryT[Entry](ctx, h, accessLogResource, accessLogSelect, values)
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

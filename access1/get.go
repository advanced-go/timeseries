package access1

import (
	"context"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []Entry, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	rows, status1 := pgxsql.Query(ctx, h, rscAccessLog, accessLogSelect, values)
	if !status1.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, h2, status1
	}
	t, status = pgxsql.Scan[Entry](rows)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, h2, status
	}
	if len(t) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

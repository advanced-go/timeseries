package access1

import (
	"context"
	"errors"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, url *url.URL) (t []Entry, status *core.Status) {
	var e E

	if url == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: URL is nil"))
	}
	rows, status1 := pgxsql.Query(ctx, h, rscAccessLog, accessLogSelect, url.Query())
	if !status1.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, status1
	}
	t, status = pgxsql.Scan[Entry](rows)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, status
	}
	if len(t) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

package accesslog

import (
	"context"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/jackc/pgx/v5"
	"net/http"
	"net/url"
)

const (
	getLoc             = PkgPath + ":get"
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
	getRouteName       = "get-entry"
	getEntryLoc        = PkgPath + ":GetEntry"
)

func getEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (t []Entry, status runtime.Status) {
	var e E

	t, status = get(ctx, h, values)
	if !status.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return
	}
	if len(t) == 0 {
		status = runtime.NewStatus(http.StatusNotFound)
	}
	return
}

func get(ctx context.Context, h http.Header, values url.Values) (t []Entry, status runtime.Status) {
	if url, override := lookup.Value(rscAccessLog); override {
		return runtime.New[[]Entry](url, nil)
	}
	var status1 runtime.Status
	var rows pgx.Rows

	rows, status1 = pgxsql.Query(ctx, h, rscAccessLog, accessLogSelect, values)
	if !status1.OK() {
		return nil, status1.AddLocation(getLoc)
	}
	t, status = pgxsql.Scan[Entry](rows)
	if !status.OK() {
		status.AddLocation(getLoc)
	}
	return
}

/*
// get - function to query for a set of entries, type selected via content Uri, from a datastore
func get(ctx context.Context, contentUri string, values map[string][]string) (any, runtime.Status) {
	rows, status := pgxsql.Query(ctx, pgxsql.NewQueryRequestFromValues(resourceNSS, accessLogSelect, values))
	if !status.OK() {
		return nil, status
	}
	switch contentUri {
	case "", CurrentVariant:
		entries, err := pgxsql.Scan[Entry](rows)
		if err != nil {
			return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		return entries, runtime.NewStatusOK() //.SetContentTypeAndLocation(CurrentVariant)
	case EntryV2Variant:
		entries, err := pgxsql.Scan[EntryV2](rows)
		if err != nil {
			return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		return entries, runtime.NewStatusOK() //.SetContentTypeAndLocation(EntryV2Variant)
	default:
		err1 := contentError(contentUri)
		return nil, runtime.NewStatusError(runtime.StatusInvalidArgument, getLoc, err1)
	}

}

func ping(ctx context.Context) runtime.Status {
	return pgxsql.Ping(ctx)
}


*/

// Scrap
/*
	switch ptr := any(&t).(type) {
	case *[]content.Entry:
		events, err := pgxsql.Scan[content.Entry](rows)
		if err != nil {
			return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		*ptr = events
	case *[]content.EntryV2:
		events, err := pgxsql.Scan[content.EntryV2](rows)
		if err != nil {
			return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		*ptr = events
	}
	return t, runtime.NewStatusOK()
*/

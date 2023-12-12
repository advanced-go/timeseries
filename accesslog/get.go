package accesslog

import (
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/postgresql/pgxsql"
	"net/http"
	"net/url"
)

const (
	getLoc             = PkgPath + ":get"
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
)

func getEntryHandler[E runtime.ErrorHandler](h http.Header, uri *url.URL) (t []Entry, status runtime.Status) {
	var e E
	//ctx := runtime.NewFileUrlContext(nil, uri.String())

	//t, status = queryEntries(ctx, uri)
	rows, status1 := pgxsql.Query(nil, pgxsql.NewQueryRequestFromValues(resourceNSS, accessLogSelect, uri.Query()))
	if !status1.OK() {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return nil, status
	}
	entries, err := pgxsql.Scan[Entry](rows)
	if err != nil {
		e.Handle(status, runtime.RequestId(h), getEntryHandlerLoc)
		return nil, runtime.NewStatusError(http.StatusInternalServerError, getEntryHandlerLoc, err)
	}
	if len(t) == 0 {
		return t, runtime.NewStatus(http.StatusNotFound)
	}
	return entries, runtime.NewStatusOK()

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

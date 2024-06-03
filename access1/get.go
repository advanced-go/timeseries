package access1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	getLoc             = PkgPath + ":get"
	getEntryHandlerLoc = PkgPath + ":getEntryHandler"
	getRouteName       = "get-entry"
	getEntryLoc        = PkgPath + ":GetEntry"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, url *url.URL) (t []Entry, status *core.Status) {
	var e E

	//t, status = get(ctx, h, url.Query())
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return
	}
	if len(t) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

/*
func get(ctx context.Context, h http.Header, values url.Values) (t []Entry, status *core.Status) {
	//if url, override := lookup.Value(rscAccessLog); override {
	//	return io2.New[[]Entry](url, nil)
	//}
	var status1 *core.Status
	var rows pgx.Rows

	rows, status1 = pgxsql.Query(ctx, h, rscAccessLog, accessLogSelect, values)
	if !status1.OK() {
		return nil, status1.AddLocation()
	}
	t, status = pgxsql.Scan[Entry](rows)
	if !status.OK() {
		status.AddLocation()
	}
	return
}


*/
/*
// get - function to query for a set of entries, type selected via content Uri, from a datastore
func get(ctx context.Context, contentUri string, values map[string][]string) (any, core.Status) {
	rows, status := pgxsql.Query(ctx, pgxsql.NewQueryRequestFromValues(resourceNSS, accessLogSelect, values))
	if !status.OK() {
		return nil, status
	}
	switch contentUri {
	case "", CurrentVariant:
		entries, err := pgxsql.Scan[Entry](rows)
		if err != nil {
			return nil, core.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		return entries, core.NewStatusOK() //.SetContentTypeAndLocation(CurrentVariant)
	case EntryV2Variant:
		entries, err := pgxsql.Scan[EntryV2](rows)
		if err != nil {
			return nil, core.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		return entries, core.NewStatusOK() //.SetContentTypeAndLocation(EntryV2Variant)
	default:
		err1 := contentError(contentUri)
		return nil, core.NewStatusError(core.StatusInvalidArgument, getLoc, err1)
	}

}

func ping(ctx context.Context) core.Status {
	return pgxsql.Ping(ctx)
}


*/

// Scrap
/*
	switch ptr := any(&t).(type) {
	case *[]content.Entry:
		events, err := pgxsql.Scan[content.Entry](rows)
		if err != nil {
			return nil, core.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		*ptr = events
	case *[]content.EntryV2:
		events, err := pgxsql.Scan[content.EntryV2](rows)
		if err != nil {
			return nil, core.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		*ptr = events
	}
	return t, core.NewStatusOK()
*/

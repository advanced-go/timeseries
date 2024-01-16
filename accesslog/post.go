package accesslog

import (
	"context"
	"errors"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/postgresql/pgxsql"
	"io"
	"net/http"
	"strings"
)

const (
	postEntryHandlerLoc = PkgPath + ":postEntryHandler"
	createEntryLoc      = PkgPath + ":createEntries"
	putLoc              = PkgPath + ":put"
	postRouteName       = "post-entry"
	postEntryLoc        = PkgPath + ":PostEntry"
)

func postEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, method string, body any) (any, runtime.Status) {
	var e E

	switch strings.ToUpper(method) {
	case http.MethodPut:
		entries, status := createEntries(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
			return nil, status
		}
		if len(entries) == 0 {
			status = runtime.NewStatusError(runtime.StatusInvalidContent, postEntryHandlerLoc, errors.New("error: no entries found"))
			e.Handle(status, runtime.RequestId(h), "")
			return nil, status
		}
		_, status = put(ctx, h, entries) // pgxsql.NewInsertRequest(h, lookup(rscAccessLog), accessLogInsert, entries[0].CreateInsertValues(entries)))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func createEntries(body any) (entries []Entry, status runtime.Status) {
	if body == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(createEntryLoc)
	}

	switch ptr := body.(type) {
	case []Entry:
		entries = ptr
	case []byte:
		entries, status = runtime.New[[]Entry](ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntryLoc)
		}
	case io.ReadCloser:
		entries, status = runtime.New[[]Entry](ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntryLoc)
		}
		//buf, status := io2.ReadAll(ptr)
		//if !status.OK() {
		//	return nil, status.AddLocation(createEntryLoc)
		//}
		//status = json2.Unmarshal(buf, &entries)
		//if !status.OK() {
		//	return nil, status.AddLocation(createEntryLoc)
		//}
	default:
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, createEntryLoc, runtime.NewInvalidBodyTypeError(body))
	}
	return entries, runtime.StatusOK()
}

// put - function to Put a set of log entries into a datastore
func put(ctx context.Context, h http.Header, entries []Entry) (tag pgxsql.CommandTag, status runtime.Status) {
	if url, override := lookup.Value(rscAccessLog); override {
		return runtime.New[pgxsql.CommandTag](url)
	}
	tag, status = pgxsql.Insert(ctx, h, rscAccessLog, accessLogInsert, entries[0].CreateInsertValues(entries))
	if !status.OK() {
		status.AddLocation(putLoc)
	}
	return
}

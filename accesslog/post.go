package accesslog

import (
	"context"
	"errors"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
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
		_, status = put(ctx, pgxsql.NewInsertRequest(h, lookup(rscAccessLog), accessLogInsert, entries[0].CreateInsertValues(entries)))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}

func createEntries(body any) ([]Entry, runtime.Status) {
	if body == nil {
		return nil, runtime.NewStatus(runtime.StatusInvalidContent).AddLocation(createEntryLoc)
	}
	var entries []Entry

	switch ptr := body.(type) {
	case []Entry:
		entries = ptr
	case []byte:
		status := json2.Unmarshal(ptr, &entries)
		if !status.OK() {
			return nil, status.AddLocation(createEntryLoc)
		}
	case io.ReadCloser:
		buf, status := io2.ReadAll(ptr)
		if !status.OK() {
			return nil, status.AddLocation(createEntryLoc)
		}
		status = json2.Unmarshal(buf, &entries)
		if !status.OK() {
			return nil, status.AddLocation(createEntryLoc)
		}
	default:
		return nil, runtime.NewStatusError(runtime.StatusInvalidContent, createEntryLoc, runtime.NewInvalidBodyTypeError(body))
	}
	return entries, runtime.StatusOK()
}

// put - function to Put a set of log entries into a datastore
func put(ctx context.Context, req pgxsql.Request) (tag pgxsql.CommandTag, status runtime.Status) {
	if req.IsFileScheme() {
		return pgxsql.CommandTag{}, io2.ReadStatus(req.Uri())
	}
	tag, status = pgxsql.Exec(ctx, req)
	if !status.OK() {
		status.AddLocation(putLoc)
	}
	return
}

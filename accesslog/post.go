package accesslog

import (
	"context"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"io"
	"net/http"
	"strings"
)

const (
	postEntryHandlerLoc = PkgPath + ":postEntryHandler"
	createEntryLoc      = PkgPath + ":createEntry"
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
		_, status = put(ctx, h, entries, "test")
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

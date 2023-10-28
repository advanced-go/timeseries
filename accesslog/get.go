package accesslog

import (
	"context"
	"encoding/json"
	"github.com/go-ai-agent/core/runtime"
	//"github.com/go-http-utils/headers"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"github.com/go-ai-agent/timeseries/accesslog/content"
)

// GetConstraints - interface defining constraints for the Get function
type GetConstraints interface {
	[]content.Entry | []content.EntryV2
}

var (
	getLoc = pkgPath + "/Get"
)

// Get - templated function to query for a set of AccessLog entries from a datastore
func Get[E runtime.ErrorHandler, T GetConstraints](ctx context.Context, values map[string][]string) (T, *runtime.Status) {
	var e E
	var t T

	rows, status := pgxsql.Query(ctx, pgxsql.NewQueryRequestFromValues(content.ResourceNSS, accessLogSelect, values))
	if !status.OK() {
		e.HandleStatus(status, ctx, getLoc)
		return nil, status
	}
	switch ptr := any(&t).(type) {
	case *[]content.Entry:
		events, err := pgxsql.Scan[content.Entry](rows)
		if err != nil {
			return nil, e.Handle(ctx, getLoc, err)
		}
		*ptr = events
	case *[]content.EntryV2:
		events, err := pgxsql.Scan[content.EntryV2](rows)
		if err != nil {
			return nil, e.Handle(ctx, getLoc, err)
		}
		*ptr = events
	}
	return t, runtime.NewStatusOK()
}

// GetByte - templated function to query for a set of AccessLog entries from a datastore
func GetByte[E runtime.ErrorHandler](ctx context.Context, contentLocation string, values map[string][]string) ([]byte, *runtime.Status) {
	var e E
	var buf []byte
	var err error

	switch contentLocation {
	case "":
		events, status := Get[E, []content.Entry](ctx, values)
		if !status.OK() {
			return nil, status
		}
		buf, err = json.Marshal(events)
	case content.Variant2Uri:
		events, status := Get[E, []content.EntryV2](ctx, values)
		if !status.OK() {
			return nil, status
		}
		buf, err = json.Marshal(events)
	default:
		err1 := contentError(contentLocation)
		return nil, e.Handle(ctx, getLoc, err1).SetCode(runtime.StatusInvalidArgument).SetContent(err1)
	}
	if err != nil {
		return nil, e.Handle(ctx, getLoc, err)
	}
	return buf, runtime.NewStatusOK() //.SetMetadata(runtime.ContentType, runtime.ContentTypeJson)
}

func ping[E runtime.ErrorHandler](ctx context.Context) *runtime.Status {
	return pgxsql.Ping[E](ctx)
}

package accesslog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/postgresql/pgxdml"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"github.com/go-ai-agent/timeseries/accesslog/content"
)

// PutConstraints - generic constraints
type PutConstraints interface {
	[]content.Entry | []content.EntryV2
}

var (
	putLoc = pkgPath + "/put"
)

func contentError(contentLocation string) error {
	return errors.New(fmt.Sprintf("invalid content location: [%v]", contentLocation))
}

// Put - templated function to Put a set of log entries into a datastore
func Put[E runtime.ErrorHandler, T PutConstraints](ctx context.Context, t T) (pgxsql.CommandTag, *runtime.Status) {
	var count = 0
	var req *pgxsql.Request

	if t == nil {
		return pgxsql.CommandTag{}, runtime.NewStatus(runtime.StatusInvalidArgument)
	}
	switch events := any(t).(type) {
	case []content.Entry:
		count = len(events)
		if count > 0 {
			req = pgxsql.NewInsertRequest(content.ResourceNSS, accessLogInsert, events[0].CreateInsertValues(events))
		}
	case []content.EntryV2:
		count = len(events)
		if count > 0 {
			req = pgxsql.NewInsertRequest(content.ResourceNSS, accessLogInsert, events[0].CreateInsertValues(events))
		}
	}
	if count > 0 {
		return pgxsql.Exec(ctx, req)
	}
	return pgxsql.CommandTag{}, runtime.NewStatusOK()
}

// PutByte - templated function to Put a set of log entries into a datastore
func PutByte[E runtime.ErrorHandler](ctx context.Context, contentLocation string, data []byte) (pgxsql.CommandTag, *runtime.Status) {
	var e E

	if data == nil {
		return pgxsql.CommandTag{}, runtime.NewStatus(runtime.StatusInvalidArgument)
	}
	switch contentLocation {
	case "":
		var events []content.Entry
		err := json.Unmarshal(data, &events)
		if err != nil {
			return pgxsql.CommandTag{}, e.Handle(ctx, putLoc, err)
		}
		return Put[E, []content.Entry](ctx, events)
	case content.Variant2Uri:
		var events []content.EntryV2
		err := json.Unmarshal(data, &events)
		if err != nil {
			return pgxsql.CommandTag{}, e.Handle(ctx, putLoc, err)
		}
		return Put[E, []content.EntryV2](ctx, events)
	default:
		err1 := contentError(contentLocation)
		return pgxsql.CommandTag{}, e.Handle(ctx, getLoc, err1).SetCode(runtime.StatusInvalidArgument).SetContent(err1)
	}
}

func delete[E runtime.ErrorHandler](ctx context.Context, where []pgxdml.Attr) (pgxsql.CommandTag, *runtime.Status) {
	if len(where) > 0 {
		return exec[E](ctx, pgxsql.NewDeleteRequest(content.ResourceNSS, deleteSql, where))
	}
	return pgxsql.CommandTag{}, runtime.NewStatusOK()
}

func exec[E runtime.ErrorHandler](ctx context.Context, req *pgxsql.Request) (pgxsql.CommandTag, *runtime.Status) {
	return pgxsql.Exec(ctx, req)
}

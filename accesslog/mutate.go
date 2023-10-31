package accesslog

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-ai-agent/core/json"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/postgresql/pgxdml"
	"github.com/go-ai-agent/postgresql/pgxsql"
)

var (
	putLoc     = pkgPath + "/put"
	putByteLoc = pkgPath + "/putByte"
)

func contentError(contentLocation string) error {
	return errors.New(fmt.Sprintf("invalid content location: [%v]", contentLocation))
}

// put - function to Put a set of log entries into a datastore
func put(ctx context.Context, contentUri string, data any) (pgxsql.CommandTag, *runtime.Status) {
	var count = 0
	var req *pgxsql.Request

	if data == nil {
		return pgxsql.CommandTag{}, runtime.NewStatus(runtime.StatusInvalidArgument)
	}
	switch contentUri {
	case "", CurrentEntryUri:
		if entries, ok := data.([]Entry); ok {
			count = len(entries)
			if count > 0 {
				req = pgxsql.NewInsertRequest(resourceNSS, accessLogInsert, entries[0].CreateInsertValues(entries))
			}
		} else {
			return pgxsql.CommandTag{}, runtime.NewStatusError(runtime.StatusInvalidArgument, putLoc, errors.New("data type is not valid for current content"))
		}
	case EntryV2Uri:
		if entries, ok := data.([]EntryV2); ok {
			count = len(entries)
			if count > 0 {
				req = pgxsql.NewInsertRequest(resourceNSS, accessLogInsert, entries[0].CreateInsertValues(entries))
			}
		} else {
			return pgxsql.CommandTag{}, runtime.NewStatusError(runtime.StatusInvalidArgument, putLoc, errors.New(fmt.Sprintf("data type is not valid for content: %v", contentUri)))
		}
	default:
		err1 := contentError(contentUri)
		return pgxsql.CommandTag{}, runtime.NewStatusError(runtime.StatusInvalidArgument, putLoc, err1)
	}
	if count > 0 {
		ct, status := pgxsql.Exec(ctx, req)
		if !status.OK() {
			return pgxsql.CommandTag{}, status
		}
		return ct, status
	}
	return pgxsql.CommandTag{}, runtime.NewStatusOK()
}

// putByte - function to Put a set of log entries into a datastore
func putByte(ctx context.Context, contentUri string, data []byte) (pgxsql.CommandTag, *runtime.Status) {
	if data == nil {
		return pgxsql.CommandTag{}, runtime.NewStatus(runtime.StatusInvalidArgument)
	}
	switch contentUri {
	case "", CurrentEntryUri:
		var events []Entry
		status := json.Unmarshal(data, &events)
		if !status.OK() {
			return pgxsql.CommandTag{}, status
		}
		return put(ctx, contentUri, data)
	case EntryV2Uri:
		var events []EntryV2
		status := json.Unmarshal(data, &events)
		if status != nil {
			return pgxsql.CommandTag{}, status
		}
		return put(ctx, contentUri, events)
	default:
		return pgxsql.CommandTag{}, runtime.NewStatusError(runtime.StatusInvalidArgument, putByteLoc, contentError(contentUri))
	}
}

func remove(ctx context.Context, where []pgxdml.Attr) (pgxsql.CommandTag, *runtime.Status) {
	if len(where) > 0 {
		return exec(ctx, pgxsql.NewDeleteRequest(resourceNSS, deleteSql, where))
	}
	return pgxsql.CommandTag{}, runtime.NewStatusOK()
}

func exec(ctx context.Context, req *pgxsql.Request) (pgxsql.CommandTag, *runtime.Status) {
	return pgxsql.Exec(ctx, req)
}

// Scrap
/*
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
*/

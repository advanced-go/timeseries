package accesslog

import (
	"context"
	"encoding/json"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"net/http"
)

var (
	getLoc = pkgPath + "/get"
)

// get - function to query for a set of entries, type selected via content Uri, from a datastore
func get(ctx context.Context, contentUri string, values map[string][]string) (any, *runtime.Status) {
	rows, status := pgxsql.Query(ctx, pgxsql.NewQueryRequestFromValues(resourceNSS, accessLogSelect, values))
	if !status.OK() {
		return nil, status
	}
	switch contentUri {
	case "", CurrentEntryUri:
		entries, err := pgxsql.Scan[Entry](rows)
		if err != nil {
			return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		return entries, runtime.NewStatusOK().SetContentTypeAndLocation(CurrentEntryUri)
	case EntryV2Uri:
		entries, err := pgxsql.Scan[EntryV2](rows)
		if err != nil {
			return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err)
		}
		return entries, runtime.NewStatusOK().SetContentTypeAndLocation(EntryV2Uri)
	default:
		err1 := contentError(contentUri)
		return nil, runtime.NewStatusError(runtime.StatusInvalidArgument, getLoc, err1)
	}

}

// getByte - templated function to query for a set of AccessLog entries from a datastore
func getByte(ctx context.Context, contentLocation string, values map[string][]string) ([]byte, *runtime.Status) {
	var buf []byte
	var err error

	switch contentLocation {
	case "", CurrentEntryUri:
		entries, status := get(ctx, contentLocation, values)
		if !status.OK() {
			return nil, status
		}
		buf, err = json.Marshal(entries)
	case EntryV2Uri:
		events, status := get(ctx, contentLocation, values)
		if !status.OK() {
			return nil, status
		}
		buf, err = json.Marshal(events)
	default:
		err1 := contentError(contentLocation)
		return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err1) //.SetCode(runtime.StatusInvalidArgument).SetContent(err1, false)
	}
	if err != nil {
		return nil, runtime.NewStatusError(http.StatusInternalServerError, getLoc, err)
	}
	return buf, runtime.NewStatusOK() //.SetMetadata(runtime.ContentType, runtime.ContentTypeJson)
}

func ping(ctx context.Context) *runtime.Status {
	return pgxsql.Ping(ctx)
}

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

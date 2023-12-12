package accesslog

import (
	"context"
	"errors"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/postgresql/pgxsql"
	"net/http"
)

const (
	putLoc = PkgPath + ":put"
)

// put - function to Put a set of log entries into a datastore
func put(ctx context.Context, h http.Header, entries []Entry, rsc string) (tag pgxsql.CommandTag, status runtime.Status) {
	if len(entries) == 0 {
		return pgxsql.CommandTag{}, runtime.NewStatusError(runtime.StatusInvalidArgument, putLoc, errors.New("entries are nil"))
	}
	req := pgxsql.NewInsertRequest(h, resolve(rsc), accessLogInsert, entries[0].CreateInsertValues(entries))
	if req.IsFileScheme() {
		return pgxsql.CommandTag{}, io2.ReadStatus(req.Uri())
	}
	tag, status = pgxsql.Exec(ctx, req)
	if !status.OK() {
		status.AddLocation(putLoc)
	}
	return
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

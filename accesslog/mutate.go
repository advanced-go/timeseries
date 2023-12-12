package accesslog

import (
	"context"
	"errors"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/postgresql/pgxsql"
)

const (
	putLoc = PkgPath + ":put"
)

// put - function to Put a set of log entries into a datastore
func put(ctx context.Context, entries []Entry) runtime.Status {
	var req pgxsql.Request

	if len(entries) == 0 {
		return runtime.NewStatusError(runtime.StatusInvalidArgument, putLoc, errors.New("entries are nil"))
	}
	req = pgxsql.NewInsertRequest(resourceNSS, accessLogInsert, entries[0].CreateInsertValues(entries))
	_, status := pgxsql.Exec(ctx, req)
	if !status.OK() {
		return status.AddLocation(putLoc)
	}
	return status
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

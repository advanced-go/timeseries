package accesslog

import (
	"context"
	"github.com/advanced-go/core/access"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"net/url"
)

type pkg struct{}

const (
	PkgPath = "github/advanced-go/timeseries/accesslog"
	//Pattern = "/" + PkgPath + "/"
	//httpHandlerRouteName = "http-handler"
	//entryResource        = "entry"

)

// GetEntry - get entries with headers and values
func GetEntry(ctx context.Context, h http.Header, values url.Values) (entries []Entry, status *runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, http.MethodGet, getEntryLoc), getRouteName, "", -1, "", access.StatusCode(&status))()
	return getEntryHandler[runtime.Log](ctx, h, values)
}

// PostEntryConstraints - Post constraints
type PostEntryConstraints interface {
	[]Entry | []byte | runtime.Nillable
}

// PostEntry - exchange function
func PostEntry[T PostEntryConstraints](ctx context.Context, h http.Header, method string, body T) (t any, status *runtime.Status) {
	h = runtime.AddRequestId(h)
	defer access.LogDeferred(access.InternalTraffic, access.NewRequest(h, method, postEntryLoc), postRouteName, "", -1, "", access.StatusCode(&status))()
	return postEntryHandler[runtime.Log](ctx, h, method, body)
}

/*
// HttpHandler - Http endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http2.WriteResponse[runtime.Log](w, nil, runtime.NewStatus(runtime.StatusInvalidArgument), nil)
		return
	}
	_, rsc, ok := uri.UprootUrn(r.URL.Path)
	if !ok || len(rsc) == 0 {
		status := runtime.NewStatusWithContent(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid path, not a valid URN: %v", r.URL.Path)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
		return
	}
	runtime.AddRequestId(r)
	switch strings.ToLower(rsc) {
	case entryResource:
		func() (status runtime.Status) {
			defer access.LogDeferred(access.InternalTraffic, r, httpHandlerRouteName, "", -1, "", &status)()
			return httpEntryHandler[runtime.Log](w, r)
		}()
	default:
		status := runtime.NewStatusWithContent(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource was not found: %v", rsc)), false)
		http2.WriteResponse[runtime.Log](w, nil, status, nil)
	}
}


*/
// Scrap
//rows, status := pgxsql.Query(rc.Context(), pgxsql.NewQueryRequestFromValues(content.ResourceNSS, accessLogSelect, values))
//if !status.OK() {
//	e.HandleStatus(status, requestId, getLoc)
//	return nil, status
//}
//return nil, runtime.NewStatusOK()

//
/*
	switch r.Method {
	case http.MethodGet:
		entries, status := GetByte[runtime.LogError](runtime.ContextWithRequest(r), httpx.GetContentLocation(r), r.URL.Query())
		if !status.OK() {
			var e E
			e.HandleStatus(status, requestId, "")
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, data, status, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson}})
		return status
	case http.MethodPut:
		var e E

		buf, status := httpx.ReadAll(r.Body)
		if !status.OK() {
			e.HandleStatus(status, nil, "")
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		if buf == nil {
			nc := runtime.NewStatus(runtime.StatusInvalidContent)
			httpx.WriteResponse[E](w, nil, nc, nil)
			return nc
		}
		_, status = PutByte[runtime.LogError](runtime.ContextWithRequest(r), httpx.GetContentLocation(r), buf)
		if !status.OK() {
			e.HandleStatus(status, "", "")
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, nil, status, nil)
	default:
	}

*/

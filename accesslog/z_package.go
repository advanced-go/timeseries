package accesslog

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/json"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

var (
	PkgUri              = reflect.TypeOf(any(pkg{})).PkgPath()
	HttpHandlerEndpoint = pkgPath + "/HttpHandler"
	EntryUri            = PkgUri + "/" + reflect.TypeOf(Entry{}).Name()
	EntryV2Uri          = PkgUri + "/" + reflect.TypeOf(EntryV2{}).Name()
	CurrentEntryUri     = EntryUri

	pkgPath        = runtime.PathFromUri(PkgUri)
	locTypeHandler = pkgPath + "/typeHandler"
	locHttpHandler = pkgPath + "/httpHandler"
	resourceNID    = "timeseries"
	resourceNSS    = "access-log"
)

// newTypeHandler - templated function providing a TypeHandlerFn via a closure
//func newTypeHandler[E runtime.ErrorHandler]() runtime.TypeHandlerFn {
//	return func(r *http.Request, body any) (any, *runtime.Status) {
//		return typeHandler[E](r, body)
//	}
//}

func CastEntry(t any) []Entry {
	if e, ok := t.([]Entry); ok {
		return e
	}
	return nil
}

func CastEntryV2(t any) []EntryV2 {
	if e, ok := t.([]EntryV2); ok {
		return e
	}
	return nil
}

// InConstraints - interface defining constraints for the Get function
type InConstraints interface {
	[]Entry | []EntryV2 | runtime.Nil
}

func TypeHandler[T InConstraints](r *http.Request, body T) (any, *runtime.Status) {
	return typeHandler[runtime.LogError](r, body)
}

func typeHandler[E runtime.ErrorHandler](r *http.Request, body any) (any, *runtime.Status) {
	var e E

	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	// create a new context with a request id. Not creating a new request as upstream processing doesn't
	// use http
	requestId := runtime.GetOrCreateRequestId(r)
	nc := runtime.ContextWithRequestId(r.Context(), requestId)
	switch r.Method {
	case http.MethodGet:
		entries, status := get(nc, r.Header.Get(httpx.ContentLocation), r.URL.Query())
		if !status.OK() {
			e.HandleStatus(status, requestId, locTypeHandler)
			return nil, status
		}
		if entries == nil {
			status.SetCode(http.StatusNotFound)
		}
		return entries, status
	case http.MethodPut:
		cmdTag, status := put(nc, r.Header.Get(httpx.ContentLocation), body)
		if !status.OK() {
			e.HandleStatus(status, requestId, locTypeHandler)
			return nil, status
		}
		return cmdTag, status
	default:
	}
	return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	httpHandler[runtime.LogError](w, r)
}

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	var e E

	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	// create a new context with a request id. Not creating a new request as upstream processing doesn't
	// use http
	requestId := runtime.GetOrCreateRequestId(r)
	nc := runtime.ContextWithRequestId(r.Context(), requestId)
	switch r.Method {
	case http.MethodGet:
		entries, status := get(nc, r.Header.Get(httpx.ContentLocation), r.URL.Query())
		if !status.OK() {
			e.HandleStatus(status, requestId, locHttpHandler)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		if entries == nil {
			status = runtime.NewStatus(http.StatusNotFound)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		var buf []byte
		buf, status = json.Marshal(entries)
		if !status.OK() {
			e.HandleStatus(status, requestId, locHttpHandler)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson},
			{httpx.ContentLocation, status.Header().Get(httpx.ContentLocation)}})
		return status
	case http.MethodPut:
		buf, status := httpx.ReadAll(r.Body)
		if !status.OK() {
			e.HandleStatus(status, requestId, locHttpHandler)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		if buf == nil {
			status = runtime.NewStatus(runtime.StatusInvalidContent)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		_, status = putByte(nc, httpx.GetContentLocation(r), buf)
		if !status.OK() {
			e.HandleStatus(status, requestId, locHttpHandler)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, nil, status, nil)
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewStatus(http.StatusMethodNotAllowed)
}

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

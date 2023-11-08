package accesslog

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/json"
	"github.com/go-ai-agent/core/log"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
)

type pkg struct{}

var (
	PkgUri         = reflect.TypeOf(any(pkg{})).PkgPath()
	Pattern        = pkgPath + "/"
	EntryVariant   = PkgUri + "/" + reflect.TypeOf(Entry{}).Name()
	EntryV2Variant = PkgUri + "/" + reflect.TypeOf(EntryV2{}).Name()
	CurrentVariant = EntryVariant

	pkgPath        = runtime.PathFromUri(PkgUri)
	locTypeHandler = pkgPath + "/typeHandler"
	locHttpHandler = pkgPath + "/httpHandler"
	//resourceNID    = "timeseries"
	resourceNSS = "access-log"
	controller  = log.NewController2(newDoHandler[runtime.LogError]())
	//typeLoc    = pkgPath + "/typeHandler"
)

// newDoHandler - templated function providing a DoHandler
func newDoHandler[E runtime.ErrorHandler]() runtime.DoHandler {
	return func(ctx any, r *http.Request, body any) (any, *runtime.Status) {
		return doHandler[E](ctx, r, body)
	}
}

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

// BodyConstraints - interface defining constraints for the TypeHandler body
type BodyConstraints interface {
	[]Entry | []EntryV2 | []byte | runtime.Nillable
}

func Do[T BodyConstraints](ctx any, method, uri, variant string, body T) (any, *runtime.Status) {
	req, status := httpx.NewRequest(ctx, method, uri, variant)
	if !status.OK() {
		return nil, status
	}
	return controller.Apply(ctx, req, body)
}

func doHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, *runtime.Status) {
	var e E

	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	switch r.Method {
	case http.MethodGet:
		entries, status := get(r.Context(), r.Header.Get(httpx.ContentLocation), r.URL.Query())
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locTypeHandler)
			return nil, status
		}
		if entries == nil {
			status.SetCode(http.StatusNotFound)
		}
		return entries, status
	case http.MethodPut:
		cmdTag, status := put(r.Context(), r.Header.Get(httpx.ContentLocation), body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locTypeHandler)
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
	r = httpx.UpdateHeadersAndContext(r)
	switch r.Method {
	case http.MethodGet:
		var buf []byte

		entries, status := Do[runtime.Nillable](r, r.Method, r.URL.String(), r.Header.Get(runtime.ContentLocation), nil)
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		buf, status = json.Marshal(entries)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locHttpHandler)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		httpx.WriteResponse[E](w, buf, status, []httpx.Attr{{httpx.ContentType, httpx.ContentTypeJson},
			{httpx.ContentLocation, status.Header().Get(httpx.ContentLocation)}})
		return status
	case http.MethodPut:
		buf, status := httpx.ReadAll(r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locHttpHandler)
			httpx.WriteResponse[E](w, nil, status, nil)
			return status
		}
		_, status = Do[[]byte](r, r.Method, r.URL.String(), r.Header.Get(runtime.ContentLocation), buf)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locHttpHandler)
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

package accesslog

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

type pkg struct{}

const (
	PkgPath        = "github.com/advanced-go/timeseries"
	Pattern        = PkgPath + "/"
	EntryVariant   = PkgPath + ":EntryV1" // + reflect.TypeOf(Entry{}).Name()
	EntryV2Variant = PkgPath + ":EntryV2" //+ reflect.TypeOf(EntryV2{}).Name()
	CurrentVariant = EntryVariant

	//pkgPath        = runtime.PathFromUri(PkgUri)
	locTypeHandler = PkgPath + "/typeHandler"
	locHttpHandler = PkgPath + "/httpHandler"
	//resourceNID    = "timeseries"
	resourceNSS = "access-log"
	//controller  = log.NewController2(newDoHandler[runtime.LogError]())
	//typeLoc    = pkgPath + "/typeHandler"
)

// newDoHandler - templated function providing a DoHandler
/*
func newDoHandler[E runtime.ErrorHandler]() runtime.DoHandler {
	return func(ctx any, r *http.Request, body any) (any, runtime.Status) {
		return doHandler[E](ctx, r, body)
	}
}
*/

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

func Do[T BodyConstraints](ctx any, method, uri, variant string, body T) (any, runtime.Status) {
	req, status := http2.NewRequest(ctx, method, uri, variant, nil)
	if !status.OK() {
		return nil, status
	}
	return doHandler[runtime.LogError](ctx, req, body)
}

func doHandler[E runtime.ErrorHandler](ctx any, r *http.Request, body any) (any, runtime.Status) {
	var e E

	if r == nil {
		return nil, runtime.NewStatus(http.StatusBadRequest)
	}
	switch r.Method {
	case http.MethodGet:
		entries, status := get(r.Context(), r.Header.Get(http2.ContentLocation), r.URL.Query())
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locTypeHandler)
			return nil, status
		}
		if entries == nil {
			status = runtime.NewStatus(http.StatusNotFound)
		}
		return entries, status
	case http.MethodPut:
		cmdTag, status := put(r.Context(), r.Header.Get(http2.ContentLocation), body)
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

func httpHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) runtime.Status {
	var e E

	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewStatus(http.StatusBadRequest)
	}
	r = http2.UpdateHeaders(r)
	switch r.Method {
	case http.MethodGet:
		var buf []byte

		entries, status := Do[runtime.Nillable](r, r.Method, r.URL.String(), r.Header.Get(http2.ContentLocation), nil)
		if !status.OK() {
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		buf, status = json2.Marshal(entries)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locHttpHandler)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, buf, status, []http2.Attr{{http2.ContentType, http2.ContentTypeJson},
			{http2.ContentLocation, status.ContentHeader().Get(http2.ContentLocation)}})
		return status
	case http.MethodPut:
		buf, status := io2.ReadAll(r.Body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locHttpHandler)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		_, status = Do[[]byte](r, r.Method, r.URL.String(), r.Header.Get(http2.ContentLocation), buf)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(r), locHttpHandler)
			http2.WriteResponse[E](w, nil, status, nil)
			return status
		}
		http2.WriteResponse[E](w, nil, status, nil)
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

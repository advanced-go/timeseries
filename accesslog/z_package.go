package accesslog

import (
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"reflect"
	"sync/atomic"
)

type pkg struct{}

var (
	pkgUri  = reflect.TypeOf(any(pkg{})).PkgPath()
	pkgPath = runtime.PathFromUri(pkgUri)

	started int64
)

// IsStarted - returns status of startup
func IsStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func setStarted() {
	atomic.StoreInt64(&started, 1)
}

func resetStarted() {
	atomic.StoreInt64(&started, 0)
}

func TimeseriesHandler(w http.ResponseWriter, r *http.Request) {
	timeseriesHandler[runtime.LogError](w, r)
}

func timeseriesHandler[E runtime.ErrorHandler](w http.ResponseWriter, r *http.Request) *runtime.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtime.NewHttpStatus(http.StatusBadRequest)
	}
	switch r.Method {
	case http.MethodGet:
		data, status := GetByte[runtime.LogError](runtime.ContextWithRequest(r), httpx.GetContentLocation(r), r.URL.Query())
		if !status.OK() {
			var e E
			e.HandleStatus(status, nil)
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		httpx.WriteResponse[E](w, data, status, httpx.ContentType, httpx.ContentTypeJson)
		return status
	case http.MethodPut:
		buf, status := httpx.ReadAll(r.Body)
		if !status.OK() {
			var e E
			e.HandleStatus(status, nil)
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		if buf == nil {
			nc := runtime.NewStatus(runtime.StatusInvalidContent)
			httpx.WriteMinResponse[E](w, nc)
			return nc
		}
		_, status = PutByte[runtime.LogError](runtime.ContextWithRequest(r), httpx.GetContentLocation(r), buf)
		if !status.OK() {
			var e E
			e.HandleStatus(status, nil)
			httpx.WriteMinResponse[E](w, status)
			return status
		}
		httpx.WriteMinResponse[E](w, status)
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return runtime.NewHttpStatus(http.StatusMethodNotAllowed)
}

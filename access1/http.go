package access1

import (
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

func httpEntryHandler[E core.ErrorHandler](w http.ResponseWriter, r *http.Request) *core.Status {
	if r == nil {
		w.WriteHeader(http.StatusBadRequest)
		return core.NewStatus(http.StatusBadRequest)
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		buf, status := get[E](r.Context(), r.Header, r.URL)
		if !status.OK() {
			httpx.WriteResponse[E](w, nil, status.Code, nil, nil)
			return status
		}
		httpx.WriteResponse[E](w, []httpx.Attr{{http2.ContentType, http2.ContentTypeJson}}, status.Code, buf, nil)
		return status
	case http.MethodPut:
		// TODO
		_, status := put[E](r.Context(), r.Header, nil) //r.Body)
		httpx.WriteResponse[E](w, nil, status.Code, nil, nil)
		return status
	default:
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return core.NewStatus(http.StatusMethodNotAllowed)
}

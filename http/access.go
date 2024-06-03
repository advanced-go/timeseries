package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/uri"
	"github.com/advanced-go/timeseries/access1"
	"github.com/advanced-go/timeseries/module"
	"net/http"
	"net/url"
)

func accessExchange(r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	if p == nil {
		p1, status := httpx.ValidateURL(r.URL, module.Authority)
		if !status.OK() {
			return httpx.NewResponseWithStatus(status, status.Err)
		}
		p = p1
	}
	r.URL = p.PathURL()
	switch r.Method {
	case http.MethodGet:
		return get(r.Context(), r.Header, r.URL, p.Version)
	case http.MethodPut:
		return put(r, p.Version)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}

func get(ctx context.Context, h http.Header, url *url.URL, version string) (resp *http.Response, status *core.Status) {
	var entries any

	switch version {
	case module.Ver1, "":
		entries, status = access1.Get(ctx, h, url)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	if !status.OK() {
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	resp, status = httpx.NewJsonResponse(entries, nil)
	if !status.OK() {
		var e core.Log
		e.Handle(status, core.RequestId(h))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	return
}

func put(r *http.Request, version string) (resp *http.Response, status *core.Status) {
	switch version {
	case module.Ver1, "":
		status = access1.Put(r, nil)
	default:
		status1 := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
		return httpx.NewResponseWithStatus(status1, status1.Err)
	}
	return httpx.NewResponseWithStatus(status, status.Err)
}

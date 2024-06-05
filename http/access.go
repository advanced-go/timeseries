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

func accessExchange[E core.ErrorHandler](r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if p == nil {
		p1, status := httpx.ValidateURL(r.URL, module.Authority)
		if !status.OK() {
			return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
		}
		p = p1
	}
	switch r.Method {
	case http.MethodGet:
		return get[E](r.Context(), r.Header, r.URL, p.Version)
	case http.MethodPut:
		return put[E](r, p.Version)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
	}
}

func get[E core.ErrorHandler](ctx context.Context, h http.Header, url *url.URL, version string) (resp *http.Response, status *core.Status) {
	var entries any
	var h2 http.Header

	switch version {
	case module.Ver1, "":
		entries, h2, status = access1.Get(ctx, h, url.Query())
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", h.Get(core.XVersion))))
		//return httpx.NewResponseWithStatus(status, status.Err)
	}
	if h2 == nil {
		h2 = make(http.Header)
	}
	if !status.OK() {
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
	}
	h2.Add(httpx.ContentType, httpx.ContentTypeJson)
	return httpx.NewResponse[E](status.HttpCode(), h2, entries)
}

func put[E core.ErrorHandler](r *http.Request, version string) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch version {
	case module.Ver1, "":
		h2, status = access1.Put(r, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	if h2 == nil {
		h2 = make(http.Header)
	}
	h2.Add(httpx.ContentType, httpx.ContentTypeText)
	return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
}

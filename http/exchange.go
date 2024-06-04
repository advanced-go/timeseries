package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/timeseries/module"
	"net/http"
	"strings"
)

// https://localhost:8081/github/advanced-go/observation:v1/search?q=golang

var (
	authorityResponse = httpx.NewAuthorityResponse(module.Authority)
)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		status := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		return httpx.NewResponse(status, status.Err), status
	}
	core.AddRequestId(r.Header)
	//r.Header.Set(core.XAuthority, module.Authority)
	switch strings.ToLower(p.Resource) {
	case module.AccessResource:
		return accessExchange(r, p)
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Resource)))
		return httpx.NewResponse(status, status.Err), status
	}
}

package accesslog

import (
	"github.com/advanced-go/core/runtime"
)

const (
	rscAccessLog = "access-log"
)

type resolverFunc func(string) string

var (
	resolverList []resolverFunc
)

func addResolver(fn resolverFunc) {
	if !runtime.IsDebugEnvironment() || fn == nil {
		return
	}
	// do not need mutex, as this is only called from test
	resolverList = append(resolverList, fn)
}

// resolve - resolve a string to an url.
func resolve(s string) string {
	if !runtime.IsDebugEnvironment() {
		return defaultResolver(s)
	}
	if resolverList != nil {
		for _, r := range resolverList {
			rsc := r(s)
			if len(rsc) != 0 {
				return rsc
			}
		}
	}
	return defaultResolver(s)
}

func defaultResolver(s string) string {
	switch s {
	default:
	}
	return s
}

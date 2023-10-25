package accesslog

import (
	"github.com/go-ai-agent/core/runtime"
	"reflect"
	"sync/atomic"
)

type pkg struct{}

var (
	PkgUrl  = runtime.ParsePkgUrl(reflect.TypeOf(any(pkg{})).PkgPath())
	PkgUri  = PkgUrl.Host + PkgUrl.Path
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

package accesslog

import (
	"github.com/advanced-go/core/uri"
)

const (
	rscAccessLog = "access-log"
)

var (
	lookup = uri.NewLookup()
)

/*
func setOverrideLookup(t any) {
	if t == nil {
		overrideLookup = nil
		return
	}
	overrideLookup = runtime.ListFromType(t)
}

func lookup(key string) []string {
	if overrideLookup == nil || len(key) == 0 {
		return nil
	}
	val := overrideLookup(key)
	if len(val) > 0 {
		return val
	}
	return nil
}


*/

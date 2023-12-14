package accesslog

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"reflect"
)

var (
	overrideLookup func(string) string
)

func setOverrideLookup(t any) {
	if t == nil {
		overrideLookup = nil
		return
	}
	overrideLookup = runtime.LookupFromType(t)
	if overrideLookup == nil {
		overrideLookup = func(key string) string {
			return fmt.Sprintf("error: invalid override Lookup type: %v", reflect.TypeOf(t))
		}
	}
}

func lookup(key string) string {
	if overrideLookup != nil {
		val := overrideLookup(key)
		if len(val) > 0 {
			return val
		}
	}
	return key
}

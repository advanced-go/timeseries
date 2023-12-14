package accesslog

import (
	"fmt"
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
	overrideLookup = LookupFromType(t)
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

func LookupFromType(t any) func(string) string {
	switch ptr := t.(type) {
	case string:
		return func(k string) string { return ptr }
	case map[string]string:
		return func(k string) string {
			v := ptr[k]
			if len(v) > 0 {
				return v
			}
			return k
		}
	case func(string) string:
		return ptr
	}
	return nil
}

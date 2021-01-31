package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func Equals(tb testing.TB, exp, act interface{}, name string) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)

		fmt.Printf("\033[31m%s:%d:\n\tname: %s\n\texp: %#v\n\tgot: %#v\033[39m\n", filepath.Base(file), line, name, exp, act)
		tb.FailNow()
	}
}

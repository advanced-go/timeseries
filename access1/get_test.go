package access1

import (
	"context"
	"fmt"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	accessV1 = "file://[cwd]/access1test/access-v1.json"
)

func testQuery[T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, values map[string][]string, args ...any) ([]T, *core.Status) {
	buf, status := io.ReadFile(accessV1)
	if !status.OK() {
		fmt.Printf("test: io.ReadFile() -> [status:%v]\n", status)
		return nil, status
	}
	entries, status1 := json2.New[[]T](buf, nil)
	if !status1.OK() {
		fmt.Printf("test: json2.New() -> [status:%v]\n", status1)
		return nil, status1
	}
	return entries, core.StatusOK()
}

func ExampleGet() {
	values := make(url.Values)
	entries, _, status := get[core.Output](nil, nil, values, testQuery[Entry])
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]

}

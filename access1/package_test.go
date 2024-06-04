package access1

import (
	"fmt"
	"net/http"
	url2 "net/url"
)

func ExampleGet() {
	h := make(http.Header)
	//h.Add(core.XAuthority, module.Authority)
	url, _ := url2.Parse("http://localhpst:8081/github/advanced-go/timeseries:access?region=*")
	entries, status := Get(nil, h, url)
	if !status.OK() {
		fmt.Printf("test: Query() -> [status:%v]\n", status)
	} else {
		//entries, status1 := pgxsql.Scan[Entry](rows)
		fmt.Printf("test: Query() -> [status:%v] [entries:%v]\n", status, len(entries))
	}

	//Output:
	//test: Query() -> [status:OK] [entries:2]

}

package access1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

const (
	accessLogState = "file://[cwd]/access1test/access-log.json"
)

func ExampleGetEntryHandler() {
	lookup.SetOverride(map[string]string{rscAccessLog: accessLogState})
	entries, _, status := get[core.Output](nil, nil, nil)

	fmt.Printf("test: getr() -> [status:%v] [entries:%v]\n", status, entries)

	//Output:
	//test: getEntryHandler() -> [status:OK] [entries:[{customer1 0001-01-01 00:00:00 +0000 UTC 450 450ms egress texas frisco loma alta timeseries-ingress 12345 timeseries 67890 urn:postgres:exec urn post postgres exec. 200 -1 flags 500 100 25 false 150 10 false false}]]

}

/*
func ExampleGetFromQuery() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()

		u, _ := url.Parse("https://google.com/search?region=oregon")
		entries, status1 := getEntryHandler[runtime.Output](nil, u.Query())
		cnt := len(entries)
		fmt.Printf("test: getEntryHandler(nil,url) -> [status:%v] [entries:%v]\n", status1, cnt)
	}

	//Output:
	//test: getEntryHandler(nil,url) -> [status:OK] [entries:2]

}


*/

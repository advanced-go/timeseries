package accesslog

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
)

const (
	accessLogState = "file://[cwd]/resource/access-log.json"
)

func ExampleGetEntryHandler() {
	setOverrideLookup(map[string][]string{rscAccessLog: {accessLogState}})
	t, status := getEntryHandler[runtime.Output](nil, nil, nil)

	fmt.Printf("test: getEntryHandler() -> [status:%v] [entries:%v]\n", status, t)

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

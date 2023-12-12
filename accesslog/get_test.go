package accesslog

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/postgresql/pgxsql"
	"net/url"
)

func pgxsqlIsStarted() bool {
	//_, status := pgxsql.TypeHandler(startup.StatusRequest, nil)
	//return status.OK()
	return false
}

func ExampleGet() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		entries, status1 := getEntryHandler[runtime.Output](nil, nil)
		fmt.Printf("test: getEntryHandler(nil,nil) -> [status:%v] [entries:%v]\n", status1, entries)

	}

	//Output:
	//test: getEntryHandler(nil,nil) -> [status:OK] [entries:[{customer7 2024-04-07 08:51:51.532388 -0500 CDT 0 450ms egress new mexico            0 0  0 100 0 false 0 0 false} {customer7 2024-04-07 08:51:51.532388 -0500 CDT 0 450ms egress new mexico            0 0  0 100 0 false 0 0 false}]]

}

func ExampleGetFromQuery() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()

		u, _ := url.Parse("https://google.com/search?region=oregon")
		entries, status1 := getEntryHandler[runtime.Output](nil, u)
		cnt := len(entries)
		fmt.Printf("test: getEntryHandler(nil,url) -> [status:%v] [entries:%v]\n", status1, cnt)
	}

	//Output:
	//test: getEntryHandler(nil,url) -> [status:OK] [entries:2]

}

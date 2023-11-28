package accesslog

import (
	"fmt"
	"github.com/advanced-go/postgresql/pgxsql"
	"net/url"
)

func pgxsqlIsStarted() bool {
	//_, status := pgxsql.TypeHandler(startup.StatusRequest, nil)
	//return status.OK()
	return false
}

func ExamplePing() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		status := ping(nil)
		fmt.Printf("test: Ping[runtimetest.DebugError](nil) -> [status:%v] [started:%v]\n", status, pgxsqlIsStarted())
	}

	//Output:
	//test: Ping(nil) -> [status:OK] [started:true]

}

func ExampleGet() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		events, status1 := get(nil, "", nil)
		fmt.Printf("test: get(nil,nil) -> [status:%v] [events:%v]\n", status1, events)

	}

	//Output:
	//test: get(nil,nil) -> [status:OK] [events:[{customer7 2024-04-07 08:51:51.532388 -0500 CDT 0 450ms egress new mexico            0 0  0 100 0 false 0 0 false} {customer7 2024-04-07 08:51:51.532388 -0500 CDT 0 450ms egress new mexico            0 0  0 100 0 false 0 0 false}]]

}

func ExampleGetFromQuery() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()

		u, _ := url.Parse("https://google.com/search?region=oregon")
		events, status1 := get(nil, "", u.Query())
		cnt := 0
		if entries, ok := events.([]Entry); ok {
			cnt = len(entries)
		}
		fmt.Printf("test: get(nil,_,url) -> [status:%v] [events:%v]\n", status1, cnt)

	}

	//Output:
	//test: get(nil,_,url) -> [status:OK] [events:2]

}

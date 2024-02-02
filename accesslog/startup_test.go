package accesslog

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/uri"
	"github.com/advanced-go/postgresql/pgxsql"
	"net/http"
)

const (
	serviceUrl = ""
)

func ExampleStartupPing() {
	r, _ := http.NewRequest("", "github/advanced-go/timeseries/accesslog:ping", nil)
	nid, rsc, ok := uri.UprootUrn(r.URL.Path)
	status := messaging.Ping(nil, nid)
	fmt.Printf("test: Ping() -> [nid:%v] [nss:%v] [ok:%v] [status-code:%v]\n", nid, rsc, ok, status.Code)

	//Output:
	//test: Ping() -> [nid:github/advanced-go/timeseries/accesslog] [nss:ping] [ok:true] [status-code:200]

}

func Example_Startup() {
	fmt.Printf("test: pgxsql.Readiness() -> %v\n", pgxsql.Readiness())
	/*
		err := testStartup()
		if err != nil {
			fmt.Printf("test: ClientStartup() -> [error:%v]\n", err)
		} else {
			defer pgxsql.ClientShutdown()
			fmt.Printf("test: ClientStartup() -> [started:%v]\n", pgxsqlIsStarted())
			time.Sleep(time.Second * 4)
		}

	*/

	//Output:
	//test: pgxsql.Readiness() -> Not Started

}

/*

func testStartup() error {
	if pgxsqlIsStarted() {
		return nil
	}
	if serviceUrl == "" {
		return errors.New("service Url is empty")
	}
	//db := core.Resource{Uri: serviceUrl}
	//err := pgxsql.ClientStartup(db, nil)
	//c <- core.Message{Event: core.StartupEvent}
	return nil
}


*/

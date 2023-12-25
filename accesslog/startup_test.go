package accesslog

import (
	"fmt"
	"github.com/advanced-go/postgresql/pgxsql"
)

const (
	serviceUrl = ""
)

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

package accesslog

import (
	"errors"
	"fmt"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"time"
)

const (
	serviceUrl = ""
)

func Example_Startup() {
	fmt.Printf("test: IsStarted() -> %v\n", pgxsql.IsStarted())
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v]\n", err)
	} else {
		defer pgxsql.ClientShutdown()
		fmt.Printf("test: ClientStartup() -> [started:%v]\n", pgxsql.IsStarted())
		time.Sleep(time.Second * 4)
	}

	//Output:
	//test: IsStarted() -> false
	//test: ClientStartup() -> [started:true]

}

func testStartup() error {
	if pgxsql.IsStarted() {
		return nil
	}
	if serviceUrl == "" {
		return errors.New("service Url is empty")
	}
	//db := host.U..DatabaseUrl{Url: serviceUrl}
	//err := pgxsql.ClientStartup(db, nil)
	//c <- host.Message{Event: host.StartupEvent}
	return nil
}

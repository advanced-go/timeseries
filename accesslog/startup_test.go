package accesslog

import (
	"errors"
	"fmt"
	"github.com/go-ai-agent/core/runtime/startup"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"time"
)

const (
	serviceUrl = ""
)

func Example_Startup() {
	fmt.Printf("test: IsStarted() -> %v\n", pgxsqlIsStarted())
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v]\n", err)
	} else {
		defer pgxsql.ClientShutdown()
		fmt.Printf("test: ClientStartup() -> [started:%v]\n", pgxsqlIsStarted())
		time.Sleep(time.Second * 4)
	}

	//Output:
	//test: IsStarted() -> false
	//test: ClientStartup() -> [started:true]

}

func testStartup() error {
	if pgxsqlIsStarted() {
		return nil
	}
	if serviceUrl == "" {
		return errors.New("service Url is empty")
	}
	db := startup.Resource{Uri: serviceUrl}
	err := pgxsql.ClientStartup(db, nil)
	c <- startup.Message{Event: startup.StartupEvent}
	return err
}

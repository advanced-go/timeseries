package accesslog

import (
	"fmt"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"github.com/go-ai-agent/timeseries/accesslog/content"
	"net/url"
)

func ExamplePing() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsql.IsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		status := ping[runtime.DebugError](nil)
		fmt.Printf("test: Ping[runtime.DebugError](nil) -> [status:%v] [started:%v]\n", status, pgxsql.IsStarted())
	}

	//Output:
	//test: Ping[runtime.DebugError](nil) -> [status:OK] [started:true]

}

func ExampleGetByteInvalidArguement() {
	events, status := GetByte[runtime.DebugError](nil, "urn:timeseries.access-log.v5", nil)
	fmt.Printf("test: GetByte[runtime.DebugError](nil,invalid,nil) -> [status:%v] [content:%v] [events:%v]\n", status, string(status.Content()), events)

	//Output:
	//[[] github.com/idiomatic-go/timeseries/accesslog/get [invalid content location: [urn:timeseries.access-log.v5]]]
	//test: GetByte[runtime.DebugError](nil,invalid,nil) -> [status:InvalidArgument] [content:invalid content location: [urn:timeseries.access-log.v5]] [events:[]]

}

func ExampleGet() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsql.IsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		events, status1 := Get[runtime.DebugError, []content.Entry](nil, nil)
		fmt.Printf("test: Get[runtime.DebugError,[]Entry](nil,nil) -> [status:%v] [events:%v]\n", status1, events)

	}

	//Output:
	//test: Get[runtime.DebugError,[]Entry](nil,nil) -> [status:OK] [events:[{customer7 2024-04-07 08:51:51.532388 -0500 CDT 0 450ms egress new mexico            0 0  0 100 0 false 0 0 false} {customer7 2024-04-07 08:51:51.532388 -0500 CDT 0 450ms egress new mexico            0 0  0 100 0 false 0 0 false}]]

}

func ExampleGetByte() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsql.IsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		buf, status := GetByte[runtime.DebugError](nil, "", nil)
		fmt.Printf("test: GetByte[runtime.DebugError](nil,current,nil) -> [status:%v] [buf:%v]\n", status, string(buf))

	}

	//Output:
	//test: GetByte[runtime.DebugError](nil,current,nil) -> [status:OK] [buf:[{"CustomerId":"customer7","StartTime":"2024-04-07T08:51:51.532388-05:00","Duration":0,"DurationString":"450ms","Traffic":"egress","Region":"new mexico","Zone":"","SubZone":"","Service":"","InstanceId":"","RouteName":"","RequestId":"","Url":"","Protocol":"","Method":"","Host":"","Path":"","StatusCode":0,"BytesSent":0,"StatusFlags":"","Timeout":0,"RateLimit":100,"RateBurst":0,"Retry":false,"RetryRateLimit":0,"RetryRateBurst":0,"Failover":false},{"CustomerId":"customer7","StartTime":"2024-04-07T08:51:51.532388-05:00","Duration":0,"DurationString":"450ms","Traffic":"egress","Region":"new mexico","Zone":"","SubZone":"","Service":"","InstanceId":"","RouteName":"","RequestId":"","Url":"","Protocol":"","Method":"","Host":"","Path":"","StatusCode":0,"BytesSent":0,"StatusFlags":"","Timeout":0,"RateLimit":100,"RateBurst":0,"Retry":false,"RetryRateLimit":0,"RetryRateBurst":0,"Failover":false}]]

}

func ExampleGetFromQuery() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsql.IsStarted())
	} else {
		defer pgxsql.ClientShutdown()

		u, _ := url.Parse("https://google.com/search?region=oregon")
		events, status1 := Get[runtime.DebugError, []content.Entry](nil, u.Query())
		fmt.Printf("test: Get[runtime.DebugError,[]Entry](nil,url) -> [status:%v] [events:%v]\n", status1, len(events))

	}

	//Output:
	//test: Get[runtime.DebugError,[]Entry](nil,url) -> [status:OK] [events:2]

}

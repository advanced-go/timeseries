package accesslog

import (
	"fmt"
	"github.com/advanced-go/postgresql/pgxdml"
	"github.com/advanced-go/postgresql/pgxsql"
	"time"
)

const (
	updateRsc = "test"
	updateSql = "UPDATE access_log"
)

var event = Entry{
	CustomerId:     "customer11",
	StartTime:      time.Now().UTC().AddDate(1, 2, 0), //ate(2023, 1, 1, 14, 12, 15, 251097, time.UTC),
	Duration:       450,
	DurationString: "450ms",
	Traffic:        "egress",
	Region:         "california",
	Zone:           "san francisco",
	SubZone:        "loma alta",
	Service:        "timeseries",
	InstanceId:     "12345",
	RouteName:      "timeseries-egress",
	RequestId:      "67890",
	Url:            "urn:postgres:exec",
	Protocol:       "urn",
	Method:         "post",
	Host:           "postgres",
	Path:           "exec.",
	StatusCode:     200,
	BytesSent:      -1,
	StatusFlags:    "flags",
	Timeout:        500,
	RateLimit:      100,
	RateBurst:      25,
	Retry:          false,
	RetryRateLimit: 150,
	RetryRateBurst: 10,
	Failover:       false,
}

var event2 = Entry{
	CustomerId: "customer12",
	StartTime:  time.Date(2023, 2, 20, 5, 45, 12, 123456, time.UTC),
	//StartTime:      time.Now().UTC(),
	Duration:       45,
	DurationString: "45ms",
	Traffic:        "ingress",
	Region:         "nevada",
	Zone:           "las vegas",
	SubZone:        "rfd #1",
	Service:        "timeseries",
	InstanceId:     "67890",
	RouteName:      "timeseries-ingress",
	RequestId:      "1234-5678-9012",
	Url:            "urn:postgres:exec",
	Protocol:       "urn",
	Method:         "post",
	Host:           "postgres",
	Path:           "exec.",
	StatusCode:     404,
	BytesSent:      -1,
	StatusFlags:    "flags of status",
	Timeout:        300,
	RateLimit:      45,
	RateBurst:      105,
	Retry:          true,
	RetryRateLimit: 250,
	RetryRateBurst: 100,
	Failover:       true,
}

func Example_put() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		events := []Entry{event, event2}
		tag, status := put(nil, "", events)
		fmt.Printf("test: put[runtimetest.DebugError,[]Entry](nil,events) -> [status:%v] [result:%v]\n", status, tag)

		//body := &httptest.ReaderCloser{Reader: bytes.NewReader(buf), Err: nil}
		//req, _ := http.NewRequest("", "www.google.com", body)
		//status = Put[runtimetest.DebugError, *http.Request](nil, RscId, req)
		//fmt.Printf("test: Put(nil,[]byte) -> [status:%v]\n", status)
	}

	//Output:
	//test: Put[runtimetest.DebugError,[]Entry](nil,events) -> [status:OK] [result:{INSERT 0 2 2 true false false false}]

}

func Example_update() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		set := []pgxdml.Attr{{"zone", "vinton"}}
		where := []pgxdml.Attr{{"region", "iowa"}}
		req := pgxsql.NewUpdateRequest(updateRsc, updateSql, set, where)

		tag, status := exec(nil, req)
		fmt.Printf("test: update(nil,req) -> [status:%v] [result:%v]\n", status, tag)

	}

	//Output:
	//test: Update(nil,req) -> [status:OK] [result:{UPDATE 6 6 false true false false}]

}

func Example_remove() {
	err := testStartup()
	if err != nil {
		fmt.Printf("test: ClientStartup() -> [error:%v] [started:%v]\n", err, pgxsqlIsStarted())
	} else {
		defer pgxsql.ClientShutdown()
		where := []pgxdml.Attr{{"region", "texas"}}

		tag, status := remove(nil, where)
		fmt.Printf("test: remove(nil,where) -> [status:%v] [result:%v]\n", status, tag)

	}

	//Output:
	//test: Delete(nil,where) -> [status:OK] [result:{DELETE 0 0 false false true false}]

}

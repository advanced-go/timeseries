package access1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	updateRsc = "test"
	updateSql = "UPDATE access_log"
	status504 = "file://[cwd]/access1test/status-504.json"
)

var event = Entry{
	StartTime:      time.Now().UTC().AddDate(1, 2, 0), //ate(2023, 1, 1, 14, 12, 15, 251097, time.UTC),
	Duration:       450,
	Traffic:        "egress",
	Region:         "california",
	Zone:           "san francisco",
	SubZone:        "loma alta",
	InstanceId:     "12345",
	Route:          "timeseries-egress",
	RequestId:      "67890",
	Url:            "urn:postgres:exec",
	Protocol:       "urn",
	Method:         "post",
	Host:           "postgres",
	Path:           "exec.",
	StatusCode:     200,
	Bytes:          -1,
	ThresholdFlags: "flags",
	Timeout:        500,
	RateLimit:      100,
	RateBurst:      25,
}

var event2 = Entry{
	StartTime: time.Date(2023, 2, 20, 5, 45, 12, 123456, time.UTC),
	//StartTime:      time.Now().UTC(),
	Duration:       45,
	Traffic:        "ingress",
	Region:         "nevada",
	Zone:           "las vegas",
	SubZone:        "rfd #1",
	InstanceId:     "67890",
	Route:          "timeseries-ingress",
	RequestId:      "1234-5678-9012",
	Url:            "urn:postgres:exec",
	Protocol:       "urn",
	Method:         "post",
	Host:           "postgres",
	Path:           "exec.",
	StatusCode:     404,
	Bytes:          -1,
	ThresholdFlags: "flags of status",
	Timeout:        300,
	RateLimit:      45,
	RateBurst:      105,
}

func Example_put() {
	lookup.SetOverride(map[string]string{rscAccessLog: status504})
	entries := []Entry{event, event2}
	_, status := put[core.Output](nil, nil, entries)
	fmt.Printf("test: put(nil,events) -> [status:%v]\n", status)

	//Output:
	//test: put(nil,events) -> [status:Timeout [status code 504 Gateway Timeout error]]

}

/*
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


*/

package access1

import (
	"fmt"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/json"
	"time"
)

type accessLogV2 struct {
	Duration string
}

var list = []Entry{
	{time.Now().UTC(), 100, access.EgressTraffic, time.Now().UTC(), "us-west", "oregon", "dc1", "www.test-host.com", "123456", "req-id", "relate-to", "HTTP/1.1", "GET", "www.google.com", "", "https://www.google.com/search?q-golang", "/search", 200, "gzip", 12345, "google-search", "primary", 500, 98.5, 10, "RL"},
	{time.Now().UTC(), 100, access.IngressTraffic, time.Now().UTC(), "us-west", "oregon", "dc1", "localhost:8081", "123456", "req-id", "relate-to", "HTTP/1.1", "GET", "github/advanced-go/search", "", "http://localhost:8081/advanced-go/search:google?q-golang", "/search", 200, "gzip", 12345, "search", "primary", 500, 100, 10, "TO"},
}

func ExampleEntry() {
	buf, status := json.Marshal(list)
	if !status.OK() {
		fmt.Printf("test: Entry{} -> [status:%v]\n", status)
	} else {
		fmt.Printf("test: Entry{} -> %v\n", string(buf))
	}

	//Output:
	//fail

}

func _ExampleScanColumnsTemplate() {
	//log := scanColumnsTemplate[AccessLog](nil)

	//fmt.Printf("test: scanColumnsTemplate[AccessLog](nil) -> %v\n", log)

	//Output:
	//fail
}

func _ExampleScannerInterface_V1() {

	//log, status := scanRowsTemplateV1[AccessLog, AccessLog](nil)
	//fmt.Printf("test: scanRowsTemplateV1() -> [status:%v] [elem:%v] [log:%v] \n", status, reflect.TypeOf(log).Elem(), log[0].DurationString)

	//Output:
	//test: scanRowsTemplateV1() -> [status:OK] [elem:timeseries.AccessLog] [log:SCAN() TEST DURATION STRING]

}

func _ExampleScannerInterface() {
	//log, status := scanRowsTemplate[accessLogV2](nil)

	//log, status := scanRowsTemplate[AccessLog](nil)
	//fmt.Printf("test: scanRowsTemplate() -> [status:%v] [elem:%v] [log:%v] \n", status, reflect.TypeOf(log).Elem(), log[0].DurationString)

	//Output:
	//test: scanRowsTemplateV1() -> [status:OK] [elem:timeseries.AccessLog] [log:SCAN() TEST DURATION STRING]

}

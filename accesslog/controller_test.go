package accesslog

import "fmt"

func ExampleLookupController() {
	key := getControllerName
	c, status := lookupController(key)

	fmt.Printf("test: LookupController(\"%v\") -> [uri:%v] [route:%v] [status:%v]\n", key, c.Uri, c.Route, status)

	//Output:
	//test: LookupController("get") -> [uri:github/advanced-go/postgresql/pgxsql] [route:timeseries-query] [status:OK]
	
}

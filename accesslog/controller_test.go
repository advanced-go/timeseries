package accesslog

import "fmt"

func ExampleLookupController() {
	key := getControllerName
	c, status := lookupController(key)

	fmt.Printf("test: LookupController(\"%v\") -> [uri:%v] [status:%v]\n", key, c.Uri, status)

	//Output:
	//test: LookupController("get") -> [uri:github/advanced-go/postgresql/pgxsql] [status:OK]

}

package accesslog

import "fmt"

func Example_PackageUri() {
	fmt.Printf("test: pkgUri -> %v\n", pkgUri)
	fmt.Printf("test: pkgPath -> %v\n", pkgPath)

	//Output:
	//test: pkgUri -> github.com/go-ai-agent/timeseries/accesslog
	//test: pkgPath -> /go-ai-agent/timeseries/accesslog

}

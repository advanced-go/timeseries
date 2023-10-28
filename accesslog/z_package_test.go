package accesslog

import "fmt"

func Example_PackageUri() {
	fmt.Printf("test: PkgUri -> %v\n", PkgUri)
	fmt.Printf("test: pkgPath -> %v\n", pkgPath)

	//Output:
	//test: PkgUri -> github.com/go-ai-agent/timeseries/accesslog
	//test: pkgPath -> /go-ai-agent/timeseries/accesslog

}

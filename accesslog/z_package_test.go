package accesslog

import "fmt"

func Example_PackageUri() {
	fmt.Printf("test: PkgUri -> %v\n", PkgUri)
	fmt.Printf("test: pkgPath -> %v\n", pkgPath)
	fmt.Printf("test: EntryLocationUri -> %v\n", EntryUri)
	fmt.Printf("test: EntryV2LocationUri -> %v\n", EntryV2Uri)

	//Output:
	//test: PkgUri -> github.com/go-ai-agent/timeseries/accesslog
	//test: pkgPath -> /go-ai-agent/timeseries/accesslog
	//test: EntryUri -> github.com/go-ai-agent/timeseries/accesslog/Entry
	//test: EntryV2Uri -> github.com/go-ai-agent/timeseries/accesslog/EntryV2

}

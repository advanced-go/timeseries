package accesslog

import "fmt"

func Example_PackageUri() {
	fmt.Printf("test: PkgUri -> %v\n", PkgUri)
	fmt.Printf("test: pkgPath -> %v\n", pkgPath)
	fmt.Printf("test: EntryVariant -> %v\n", EntryVariant)
	fmt.Printf("test: EntryV2Variant -> %v\n", EntryV2Variant)

	//Output:
	//test: PkgUri -> github.com/go-ai-agent/timeseries/accesslog
	//test: pkgPath -> /go-ai-agent/timeseries/accesslog
	//test: EntryVariant -> github.com/go-ai-agent/timeseries/accesslog/Entry
	//test: EntryV2Variant -> github.com/go-ai-agent/timeseries/accesslog/EntryV2

}

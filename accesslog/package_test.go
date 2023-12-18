package accesslog

import "fmt"

func Example_PackageUri() {
	fmt.Printf("test: PkgPath =  \"%v\"\n", PkgPath)
	//fmt.Printf("test: pkgPath -> %v\n", pkgPath)
	//fmt.Printf("test: EntryVariant -> %v\n", EntryVariant)
	//fmt.Printf("test: EntryV2Variant -> %v\n", EntryV2Variant)

	//Output:
	//test: PkgPath =  "github.com/advanced-go/timeseries/accesslog"

}

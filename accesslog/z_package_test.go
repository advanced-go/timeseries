package accesslog

import "fmt"

func Example_PackageUri() {
	fmt.Printf("test: pkgUri -> %v\n", pkgUri)
	fmt.Printf("test: pkgPath -> %v\n", pkgPath)

	//Output:
	//test: PkgUrl -> file://github.com/go-ai-agent/postgresql/pgxsql
	//test: PkgUri -> github.com/go-ai-agent/postgresql/pgxsql
	//test: PkgPath -> /go-ai-agent/postgresql/pgxsql

}

# timeseries


## accesslog

[AccessLog][timeseriespkg] provides templated Get and Put functions for access log timeseries information. The functions use [PostgreSQL][postgresqlpkg]
to interact with TimescaleDB.

~~~
// GetConstraints - interface defining constraints for the Get function
type GetConstraints interface {
	[]content.Entry | []content.EntryV2
}

// Get - templated function to query for a set of AccessLog entries from a datastore
func Get[E runtime.ErrorHandler, T GetConstraints](ctx context.Context, values map[string][]string) (T, *runtime.Status) {
    // implementation details
}

// GetByte - templated function to query for a set of AccessLog entries from a datastore
func GetByte[E runtime.ErrorHandler](ctx context.Context, contentLocation string, values map[string][]string) ([]byte, *runtime.Status) {
    // implementation details
}

// PutConstraints - generic constraints
type PutConstraints interface {
	[]content.Entry | []content.EntryV2
}

// Put - templated function to Put a set of log entries into a datastore
func Put[E runtime.ErrorHandler, T PutConstraints](ctx context.Context, t T) (pgxsql.CommandTag, *runtime.Status) {
    // implementation details
}

// PutByte - templated function to Put a set of log entries into a datastore
func PutByte[E runtime.ErrorHandler](ctx context.Context, contentLocation string, data []byte) (pgxsql.CommandTag, *runtime.Status) {
    // implementation details
}

~~~

[timeseriespkg]: <https://pkg.go.dev/github.com/gotemplates/timeseries/accesslog>
[postgresqlpkg]: <https://pkg.go.dev/github.com/gotemplates/postgresql/pgxsql>

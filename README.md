# go/templates/timeseries


## accesslog

[AccessLog][timeseriespkg] implements types that build SQL statements based on the configured attributes. Support is also available for selecting
PostgreSQL functions for timestamps and next values.

~~~
// ExpandSelect - given a template, expand the template to build a WHERE clause if configured
func ExpandSelect(template string, where []Attr) (string, error) {
}

// WriteInsert - build a SQL insert statement with a VALUES list
func WriteInsert(sql string, values [][]any) (string, error) {
}

// WriteUpdate - build a SQL update statement, including SET and WHERE clauses
func WriteUpdate(sql string, attrs []Attr, where []Attr) (string, error) {
}

// WriteDelete - build a SQL delete statement with a WHERE clause
func WriteDelete(sql string, where []Attr) (string, error) {
}
~~~


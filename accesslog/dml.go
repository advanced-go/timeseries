package accesslog

const (
	CustomerIdName     = "customer_id"
	StartTimeName      = "start_time"
	DurationName       = "duration_ms"
	DurationStrName    = "duration_str"
	TrafficName        = "traffic"
	RegionName         = "region"
	ZoneName           = "zone"
	SubZoneName        = "sub_zone"
	ServiceName        = "service"
	InstanceIdName     = "instance_id"
	RouteNameName      = "route_name"
	RequestIdName      = "request_id"
	UrlName            = "url"
	ProtocolName       = "protocol"
	MethodName         = "method"
	HostName           = "host"
	PathName           = "path"
	StatusCodeName     = "status_code"
	BytesSentName      = "bytes_sent"
	StatusFlagsName    = "status_flags"
	TimeoutName        = "timeout"
	RateLimitName      = "rate_limit"
	RateBurstName      = "rate_burst"
	RetryName          = "retry"
	RetryRateLimitName = "retry_rate_limit"
	RetryRateBurstName = "retry_rate_burst"
	FailoverName       = "failover"
	ProxyName          = "proxy"

	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	accessLogSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"

	accessLogInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst,retry,retry_rate_limit,retry_rate_burst,failover) VALUES"

	deleteSql = "DELETE FROM access_log"
)

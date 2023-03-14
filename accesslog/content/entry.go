package content

import (
	"errors"
	"fmt"
	"github.com/gotemplates/timeseries/accesslog"
	"time"
)

type EntryV2 struct {
	CustomerId     string
	StartTime      time.Time
	Duration       int64
	DurationString string
	Traffic        string

	Region     string
	Zone       string
	SubZone    string
	Service    string
	InstanceId string
	RouteName  string

	RequestId string
	Url       string
	Protocol  string
	Method    string
	Host      string
	Path      string

	StatusCode  int32
	BytesSent   int64
	StatusFlags string

	Timeout        int32
	RateLimit      float64
	RateBurst      int32
	Retry          bool
	RetryRateLimit float64
	RetryRateBurst int32
	Failover       bool
	Proxy          bool
}

func (EntryV2) Scan(columnNames []string, values []any) (log EntryV2, err error) {
	return EntryV2{}, nil
}

func (a EntryV2) Values() []any {
	return []any{
		a.CustomerId,
		a.StartTime,
		a.Duration,
		a.DurationString,
		a.Traffic,

		a.Region,
		a.Zone,
		a.SubZone,
		a.Service,
		a.InstanceId,
		a.RouteName,

		a.RequestId,
		a.Url,
		a.Protocol,
		a.Method,
		a.Host,
		a.Path,

		a.StatusCode,
		a.BytesSent,
		a.StatusFlags,

		a.Timeout,
		a.RateLimit,
		a.RateBurst,
		a.Retry,
		a.RetryRateLimit,
		a.RetryRateBurst,
		a.Failover,
		a.Proxy,
	}
}

func (EntryV2) CreateInsertValues(events []EntryV2) [][]any {
	var values [][]any

	for _, e := range events {
		values = append(values, e.Values())
	}
	return values
}

// Entry - timeseries access log struct
type Entry struct {
	CustomerId     string
	StartTime      time.Time
	Duration       int64
	DurationString string
	Traffic        string

	Region     string
	Zone       string
	SubZone    string
	Service    string
	InstanceId string
	RouteName  string

	RequestId string
	Url       string
	Protocol  string
	Method    string
	Host      string
	Path      string

	StatusCode  int32
	BytesSent   int64
	StatusFlags string

	Timeout        int32
	RateLimit      float64
	RateBurst      int32
	Retry          bool
	RetryRateLimit float64
	RetryRateBurst int32
	Failover       bool
	Proxy          bool
}

func (Entry) Scan(columnNames []string, values []any) (log Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case accesslog.CustomerIdName:
			log.CustomerId = values[i].(string)
		case accesslog.StartTimeName:
			log.StartTime = values[i].(time.Time)
		case accesslog.DurationName:
			log.Duration = values[i].(int64)
		case accesslog.DurationStrName:
			log.DurationString = values[i].(string)
		case accesslog.TrafficName:
			log.Traffic = values[i].(string)
		case accesslog.RegionName:
			log.Region = values[i].(string)
		case accesslog.ZoneName:
			log.Zone = values[i].(string)
		case accesslog.SubZoneName:
			log.SubZone = values[i].(string)
		case accesslog.ServiceName:
			log.Service = values[i].(string)
		case accesslog.InstanceIdName:
			log.InstanceId = values[i].(string)
		case accesslog.RouteNameName:
			log.RouteName = values[i].(string)
		case accesslog.RequestIdName:
			log.RequestId = values[i].(string)
		case accesslog.UrlName:
			log.Url = values[i].(string)
		case accesslog.ProtocolName:
			log.Protocol = values[i].(string)
		case accesslog.MethodName:
			log.Method = values[i].(string)
		case accesslog.HostName:
			log.Host = values[i].(string)
		case accesslog.PathName:
			log.Path = values[i].(string)
		case accesslog.StatusCodeName:
			log.StatusCode = values[i].(int32)
		case accesslog.BytesSentName:
			log.BytesSent = values[i].(int64)
		case accesslog.StatusFlagsName:
			log.StatusFlags = values[i].(string)
		case accesslog.TimeoutName:
			log.Timeout = values[i].(int32)
		case accesslog.RateLimitName:
			log.RateLimit = values[i].(float64)
		case accesslog.RateBurstName:
			log.RateBurst = values[i].(int32)
		case accesslog.RetryName:
			log.Retry = values[i].(bool)
		case accesslog.RetryRateLimitName:
			log.RetryRateLimit = values[i].(float64)
		case accesslog.RetryRateBurstName:
			log.RetryRateBurst = values[i].(int32)
		case accesslog.FailoverName:
			log.Failover = values[i].(bool)
		case accesslog.ProxyName:
			log.Proxy = values[i].(bool)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (a Entry) Values() []any {
	return []any{
		a.CustomerId,
		a.StartTime,
		a.Duration,
		a.DurationString,
		a.Traffic,

		a.Region,
		a.Zone,
		a.SubZone,
		a.Service,
		a.InstanceId,
		a.RouteName,

		a.RequestId,
		a.Url,
		a.Protocol,
		a.Method,
		a.Host,
		a.Path,

		a.StatusCode,
		a.BytesSent,
		a.StatusFlags,

		a.Timeout,
		a.RateLimit,
		a.RateBurst,
		a.Retry,
		a.RetryRateLimit,
		a.RetryRateBurst,
		a.Failover,
		a.Proxy,
	}
}

func (Entry) CreateInsertValues(events []Entry) [][]any {
	var values [][]any

	for _, e := range events {
		values = append(values, e.Values())
	}
	return values
}

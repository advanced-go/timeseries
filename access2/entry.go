package access2

import (
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

package activity

import "time"

// Origin - attributes that uniquely identify a service instance
type Origin struct {
	Region     string
	Zone       string
	Service    string
	InstanceId string
}

type Metrics struct {
	Watch   int16
	Percent int16 // Used for latency, traffic, status codes, counter, profile
	Value   int16 // Used for latency, saturation duration or traffic
	Minimum int16 // Used for status codes to attenuate underflow, applied to the window interval
}

type Entry struct {
	CustomerId int32
	SloEntryId int32

	Timestamp time.Time

	Origin Origin

	Event     int32   // event : warning, watch, canceled
	Status    int32   // Status : warning, watch, canceled, ok/normal/canary?
	Threshold Metrics // This can't be pulled from Slo entry as this is historical
	Actual    Metrics

	LocalityScope int16 // Values : None, Region, Zone, Default - uses all localities

	Advice string // Given an event is generated, what information is the downstream service going
	// to need to know to affect a change

}

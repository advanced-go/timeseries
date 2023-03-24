package activity

import (
	"errors"
	"github.com/go-sre/core/runtime"
	"github.com/go-sre/host/messaging"
	"github.com/go-sre/postgresql/pgxsql"
	"reflect"
	"sync/atomic"
	"time"
)

type pkg struct{}

var (
	Uri      = pkgPath
	c        = make(chan messaging.Message, 1)
	pkgPath  = reflect.TypeOf(any(pkg{})).PkgPath()
	started  int64
	duration = time.Second * 4
	location = pkgPath + "/startup"
)

// IsStarted - determine if this package has completed startup
func IsStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func setStarted() {
	atomic.StoreInt64(&started, 1)
}

func init() {
	messaging.RegisterResource(Uri, c)
	go receive()
}

var messageHandler messaging.MessageHandler = func(msg messaging.Message) {
	start := time.Now()
	switch msg.Event {
	case messaging.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			if pgxsql.IsStarted() {
				messaging.ReplyTo(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
				setStarted()
				return
			}
			time.Sleep(wait)
		}
		messaging.ReplyTo(msg, runtime.NewStatusError(location, errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)))
	case messaging.ShutdownEvent:
	}
}

func receive() {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			go messageHandler(msg)
		default:
		}
	}
}

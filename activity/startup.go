package activity

import (
	"errors"
	"github.com/go-ai-agent/core/host"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"reflect"
	"sync/atomic"
	"time"
)

type pkg struct{}

var (
	Uri      = pkgPath
	c        = make(chan host.Message, 1)
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
	host.Register(Uri, c)
	go receive()
}

var messageHandler host.MessageHandler = func(msg host.Message) {
	start := time.Now()
	switch msg.Event {
	case host.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			if pgxsql.IsStarted() {
				host.ReplyTo(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
				setStarted()
				return
			}
			time.Sleep(wait)
		}
		host.ReplyTo(msg, runtime.NewStatusError(errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)).SetLocation(location))
	case host.ShutdownEvent:
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

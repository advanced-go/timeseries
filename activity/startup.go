package activity

import (
	"errors"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/messaging/core"
	"github.com/advanced-go/messaging/exchange"
	"reflect"
	"sync/atomic"
	"time"
)

type pkg struct{}

var (
	Uri      = pkgPath
	c        = make(chan core.Message, 1)
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
	exchange.Register(Uri, c)
	go receive()
}

var messageHandler core.MessageHandler = func(msg core.Message) {
	start := time.Now()
	switch msg.Event {
	case core.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			//if pgxsql.IsStarted() {
			//	startup.ReplyTo(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
			//	setStarted()
			//	return
			//}
			time.Sleep(wait)
		}
		core.ReplyTo(msg, runtime.NewStatusError(0, location, errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)))
	case core.ShutdownEvent:
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

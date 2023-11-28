package accesslog

import (
	"errors"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/messaging/core"
	"github.com/advanced-go/messaging/exchange"
	"sync/atomic"
	"time"
)

const (
	location = PkgPath + ":startup"
)

var (
	c        = make(chan core.Message, 1)
	started  int64
	duration = time.Second * 4
)

// IsStarted - determine if this package has completed startup
func isStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func setStarted() {
	atomic.StoreInt64(&started, 1)
}

func init() {
	exchange.Register(PkgPath, c)
	go receive()
}

var messageHandler core.MessageHandler = func(msg core.Message) {
	start := time.Now()
	switch msg.Event {
	case core.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			status := runtime.NewStatusOK() //pgxsql.TypeHandler(startup.StatusRequest, nil)
			if status.OK() {
				core.ReplyTo(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
				setStarted()
				return
			}
			time.Sleep(wait)
		}
		core.ReplyTo(msg, runtime.NewStatusError(runtime.StatusInvalidArgument, location, errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)))
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

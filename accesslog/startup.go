package accesslog

import (
	"errors"
	"github.com/go-ai-agent/core/resiliency"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/core/runtime/startup"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"sync/atomic"
	"time"
)

var (
	c        = make(chan startup.Message, 1)
	started  int64
	duration = time.Second * 4
	location = pkgPath + "/startup"
)

// IsStarted - determine if this package has completed startup
func isStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func setStarted() {
	atomic.StoreInt64(&started, 1)
}

func init() {
	startup.Register(PkgUri, c)
	go receive()
}

var messageHandler startup.MessageHandler = func(msg startup.Message) {
	start := time.Now()
	switch msg.Event {
	case startup.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			_, status := pgxsql.TypeHandler(startup.StatusRequest, nil)
			if status.OK() {
				startup.ReplyTo(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
				setStarted()
				return
			}
			time.Sleep(wait)
		}
		startup.ReplyTo(msg, runtime.NewStatusError(runtime.StatusInvalidArgument, location, errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)))
		initController(msg)
	case startup.ShutdownEvent:
	}
}

func initController(msg startup.Message) {
	// if a controller is configured
	var cfg resiliency.ControllerConfig
	cfg.Name = "timeseries:controller"
	cfg.Primary.Select = func(status *runtime.Status) bool { return true }
	//	cfg.
	//	controller = resiliency.NewController[runtime.LogError]()

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

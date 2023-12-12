package accesslog

import (
	"errors"
	"fmt"
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
	started  int64
	duration = time.Second * 4
	agent    exchange.Agent
)

// IsStarted - determine if this package has completed startup
func isStarted() bool {
	return atomic.LoadInt64(&started) != 0
}

func setStarted() {
	atomic.StoreInt64(&started, 1)
}

func init() {
	status := exchange.Register(exchange.NewMailbox(PkgPath, false))
	if status.OK() {
		agent, status = exchange.NewAgent(PkgPath, messageHandler, nil, nil)
	}
	if !status.OK() {
		fmt.Printf("init() failure: [%v]\n", PkgPath)
	}
	agent.Run()
}

var messageHandler core.MessageHandler = func(msg core.Message) {
	start := time.Now()
	switch msg.Event {
	case core.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			status := runtime.NewStatusOK() //pgxsql.TypeHandler(startup.StatusRequest, nil)
			if status.OK() {
				core.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
				setStarted()
				return
			}
			time.Sleep(wait)
		}
		core.SendReply(msg, runtime.NewStatusError(runtime.StatusInvalidArgument, location, errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)))
	case core.ShutdownEvent:
	}
}

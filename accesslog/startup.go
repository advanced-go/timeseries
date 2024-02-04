package accesslog

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/runtime"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	location = PkgPath + ":startup"
)

var (
	ready    int64
	duration = time.Second * 4
	agent    *messaging.Agent
)

func init() {
	var err error
	agent, err = messaging.NewDefaultAgent(PkgPath, messageHandler, false)
	if err != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err)
	}
	agent.Run()
}

var messageHandler messaging.MessageHandler = func(msg messaging.Message) {
	start := time.Now()
	switch msg.Event {
	case messaging.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			// TO DO : uncomment call to pgxsql.Readiness()
			//status := pgxsql.Readiness()
			status := runtime.StatusOK()
			if status.OK() {
				messaging.SendReply(msg, messaging.NewStatusDuration(http.StatusOK, time.Since(start)))
				setReady()
				return
			}
			time.Sleep(wait)
		}
		messaging.SendReply(msg, messaging.NewStatusDurationError(runtime.StatusInvalidArgument, time.Since(start), errors.New("startup error: pgxsql not started")))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		messaging.SendReply(msg, messaging.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}

func isReady() bool {
	return atomic.LoadInt64(&ready) != 0
}

func setReady() {
	atomic.StoreInt64(&ready, 1)
}

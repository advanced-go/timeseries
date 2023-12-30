package accesslog

import (
	"errors"
	"fmt"
	"github.com/advanced-go/core/messaging"
	"github.com/advanced-go/core/runtime"
	"sync/atomic"
	"time"
)

const (
	location = PkgPath + ":startup"
)

var (
	ready    int64
	duration = time.Second * 4
	agent    messaging.Agent
)

func isReady() bool {
	return atomic.LoadInt64(&ready) != 0
}

func setReady() {
	atomic.StoreInt64(&ready, 1)
}

func init() {
	var status runtime.Status
	agent, status = messaging.NewDefaultAgent(PkgPath, messageHandler, false)
	if !status.OK() {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, status)
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
			status := runtime.NewStatusOK()
			if status.OK() {
				messaging.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
				setReady()
				return
			}
			time.Sleep(wait)
		}
		messaging.SendReply(msg, runtime.NewStatusError(runtime.StatusInvalidArgument, location, errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		messaging.SendReply(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
	}
}

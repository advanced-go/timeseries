package accesslog

import (
	"errors"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/core/runtime/startup"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"time"
)

var (
	c        = make(chan startup.Message, 1)
	duration = time.Second * 4
	location = pkgPath + "/startup"
)

var messageHandler startup.MessageHandler = func(msg startup.Message) {
	start := time.Now()
	switch msg.Event {
	case startup.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			if pgxsql.IsStarted() {
				startup.ReplyTo(msg, runtime.NewStatusOK().SetDuration(time.Since(start)))
				setStarted()
				return
			}
			time.Sleep(wait)
		}
		startup.ReplyTo(msg, runtime.NewStatusError(0, location, errors.New("startup error: pgxsql not started")).SetDuration(time.Since(start)))
	case startup.ShutdownEvent:
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

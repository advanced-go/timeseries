package accesslog

import (
	"errors"
	"github.com/go-ai-agent/core/host"
	"github.com/go-ai-agent/core/runtime"
	"github.com/go-ai-agent/postgresql/pgxsql"
	"reflect"
	"time"
)

var (
	Uri      = pkgPath
	c        = make(chan host.Message, 1)
	pkgPath  = reflect.TypeOf(any(pkg{})).PkgPath()
	duration = time.Second * 4
	location = pkgPath + "/startup"
)

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

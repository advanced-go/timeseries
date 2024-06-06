package access1

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/host"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	ready    int64
	duration = time.Second * 4
	agent    *messaging.Agent
)

func init() {
	a, err1 := host.RegisterControlAgent(PkgPath, messageHandler)
	if err1 != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err1)
	}
	a.Run()
}

func messageHandler(msg *messaging.Message) {
	start := time.Now()
	switch msg.Event() {
	case messaging.StartupEvent:
		for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
			// TO DO : uncomment call to pgxsql.Readiness()
			//status := pgxsql.Readiness()
			status := runtime.StatusOK()
			if status.OK() {
				messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
				setReady()
				return
			}
			time.Sleep(wait)
		} // TODO
		//messaging.SendReply(msg, core.NewStatusDurationError(runtime.StatusInvalidArgument, time.Since(start), errors.New("startup error: pgxsql not started")))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}

func isReady() bool {
	return atomic.LoadInt64(&ready) != 0
}

func setReady() {
	atomic.StoreInt64(&ready, 1)
}

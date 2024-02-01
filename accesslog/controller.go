package accesslog

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"github.com/advanced-go/core/controller"
	"github.com/advanced-go/core/runtime"
	"io/fs"
	"net/http"
)

const (
	controllersPath   = "resource/controllers.json"
	controllerLookup  = PkgPath + ":LookupController"
	getControllerName = "get"
	putControllerName = "put"
)

var (
	//go:embed resource/*
	f embed.FS

	cm *controller.Map
)

func init() {
	var status runtime.Status

	buf, err := fs.ReadFile(f, controllersPath)
	if err != nil {
		fmt.Printf("controller.init(\"%v\") failure: [%v]\n", PkgPath, err)
		return
	}
	cm, status = controller.NewMap(buf)
	if !status.OK() {
		fmt.Printf("controller.init(\"%v\") failure: [%v]\n", PkgPath, status)
	}
}

func lookupController(key string) (*controller.Controller, runtime.Status) {
	if cm == nil {
		return nil, runtime.NewStatusError(runtime.StatusInvalidArgument, controllerLookup, errors.New("error: controller map is nil"))
	}
	return cm.Get(key)
}

func apply(ctx context.Context, newCtx *context.Context, controllerName, rscName string, h http.Header, statusCode controller.StatusCodeFunc) func() {
	c, _ := lookupController(controllerName)
	if c == nil {
		c = new(controller.Controller)
		c.Name = "error"
		c.Route = "error"
		c.Method = "invalid"
		c.Uri = ""
		c.Duration = 0
	}
	if len(rscName) > 0 {
		c.Uri += ":" + rscName
	}
	return controller.Apply(ctx, newCtx, c.Method, c.Uri, c.Route, h, c.Duration, statusCode)
}

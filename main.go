package main

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/pkg/drives"
	"github.com/kekaiwang/go-blog/router"
	"github.com/kekaiwang/go-blog/utils/core"
	"github.com/kekaiwang/go-blog/utils/format"
)

// main
func main() {
	initResource()
	initResources()
	wg := sync.WaitGroup{}

	// init context
	stopCtx, stopFunc := context.WithCancel(context.Background())
	webServer(stopCtx, &wg)

	// grace stop
	core.RegisterSignal(stopFunc)

	wg.Wait()
}

// web server
func webServer(stopCtx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	// init gin
	r := gin.Default()

	router.SetupRouter(r)
	ser := &http.Server{
		Addr:              ":8089",
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
	}

	// print info
	format.PrintGreen("server start : http://127.0.0.1:8089")

	go func(stopFunc context.Context, waitCtx *sync.WaitGroup) {
		// server start
		if err := gracehttp.Serve(ser); err != nil {
			panic(err.Error())
		}

		// receive stop signal
		for {
			select {
			case <-stopFunc.Done():
				format.PrintGreen("web server stopped")
				waitCtx.Done()
				return
			}
		}
	}(stopCtx, wg)
}

// init source
func initResource() {
	// db init
	drives.BlogDBInit()
}

// init source info
func initResources() {
	// db init
	drives.BlogDBInit()
}

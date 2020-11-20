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

func main() {
	initResource()
	wg := sync.WaitGroup{}

	stopCtx, stopFunc := context.WithCancel(context.Background())
	webServer(stopCtx, &wg)

	core.RegisterSignal(stopFunc)

	wg.Wait()
}

func webServer(stopCtx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	r := gin.Default()

	router.SetupRouter(r)
	ser := &http.Server{
		Addr:              ":8088",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	go func(stopFunc context.Context, waitCtx *sync.WaitGroup) {
		if err := gracehttp.Serve(ser); err != nil {
			panic(err.Error())
		}

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

func initResource() {
	drives.BlogDBInit()
}

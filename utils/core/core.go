/**
 * Author: Kekai Wang
 */
package core

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kekaiwang/go-blog/utils/format"
)

func RegisterSignal(stopFunc context.CancelFunc) {
	ch := make(chan os.Signal, 2)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func(stopFunc context.CancelFunc) {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			format.PrintGreen("receive stop signal")
			stopFunc()
			return
		}
	}(stopFunc)
}

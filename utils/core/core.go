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

// RegisterSignal.
func RegisterSignal(stopFunc context.CancelFunc) {
	ch := make(chan os.Signal, 2)

	// notify signal
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func(stopFunc context.CancelFunc) {
		sig := <-ch // send
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM: // ginal
			format.PrintGreen("receive stop signal")
			stopFunc() // stop server
			return
		}
	}(stopFunc)
}

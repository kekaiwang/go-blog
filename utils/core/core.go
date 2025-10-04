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
	ch := make(chan os.Signal, 2) // chan

	// notify signal
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM) // ginal

	go func(stopFunc context.CancelFunc) { // stop sginal
		sig := <-ch // send
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM: // ginal stop
			format.PrintGreen("receive stop signal") // print signal
			stopFunc()                               // stop server
			return
		}
	}(stopFunc)
}

// RegisterSignal.
func RegisterSignalInfo(stopFunc context.CancelFunc) {
	ch := make(chan os.Signal, 2) // chan

	// notify signal
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM) // ginal

	go func(stopFunc context.CancelFunc) { // stop sginal
		sig := <-ch // send
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM: // ginal stop
			format.PrintGreen("receive stop signal") // print signal
			stopFunc()                               // stop server
			return
		}
	}(stopFunc)
}

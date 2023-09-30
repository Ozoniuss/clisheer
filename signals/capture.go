package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var Capture = make(chan os.Signal, 1)

// ListenForTermination is listening for a ctrl+C signal in order to display a
// custom message when an interactive prompt is stopped.
func ListenForTermination() {
	signal.Notify(Capture, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-Capture
		fmt.Println("\nexited")
		os.Exit(1)
	}()
}

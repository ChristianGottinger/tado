package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func PanicOnError(err error) {
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
}

func Wait() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	<-c

	// The signal is received, you can now do the cleanup
	fmt.Print("Shutting down")
}

package main

import (
	copshq "github.com/conplementag/cops-hq/v2/pkg/hq"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	defer errorhandler()

	hq := copshq.NewCustom("tado", "0.0.1", &copshq.HqOptions{
		Quiet:              false,
		LogFileName:        "",
		DisableFileLogging: true,
	})

	createCommands(hq)

	hq.Run()
}

func errorhandler() {
	if r := recover(); r != nil {
		logrus.Errorf("Unhandled exception terminating the application: %+v\n", r)
		os.Exit(1)
	}
}

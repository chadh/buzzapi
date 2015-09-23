package main

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

const formatStr string = "%{color}%{time:2006-01-02 15:04:05} [%{pid}] %{level}:%{color:reset} %{message}"

var log = logging.MustGetLogger("")

func initializeLogger() {
	// Logging stuff
	var format = logging.MustStringFormatter(formatStr)
	logging.SetFormatter(format)

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	b2f := logging.NewBackendFormatter(backend, format)
	lb := logging.AddModuleLevel(b2f)
	logging.SetBackend(lb)
	if *verbose {
		fmt.Println("verbose mode")
		logging.SetLevel(logging.DEBUG, "")
	} else if *quiet {
		logging.SetLevel(logging.CRITICAL, "")
	} else {
		logging.SetLevel(logging.WARNING, "")
	}
}

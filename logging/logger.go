package logging

import (
	"os"

	log "github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var verboseMode = false

// InitLogger initializes the logging system
func InitLogger() {
	log.SetOutput(os.Stdout)
	prefixedFormatter := &prefixed.TextFormatter{TimestampFormat: "15:04:05.000"}

	log.SetFormatter(prefixedFormatter)
	log.SetLevel(log.InfoLevel)
}

// VerboseMode activates the verbose mode, allowing more traces to be outputted
func VerboseMode() {
	verboseMode = true
}

// DebugMode activate the debug mode, lowering the log level to print out debug traces
func DebugMode() {
	log.SetLevel(log.DebugLevel)
}

// Info prints out a Info level message
func Info(message string) {
	log.Info(message)
}

// Trace prints out a Info level message but only when in Verbose mode
func Trace(message string) {
	if verboseMode == true {
		log.Info(message)
	}
}

// Debug prints out a Debug level message
func Debug(message string) {
	log.Debug(message)
}

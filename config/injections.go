package config

import (
	"github.com/tahitianstud/utils/logging"
	"github.com/tahitianstud/utils/logging/custom"
)

// Inject loads the dependency injection configuration
func Inject() {}

// provides Dependency Injection using factory methods
func init() {
	// use the CustomLogger Implementation for our logging
	logging.New = custom.Logger
}

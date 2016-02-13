package config

import (
	"github.com/tahitianstud/utils/logging"
	"github.com/tahitianstud/utils/logging/framework"
)

var Logger = logging.New()

// Inject loads the dependency injection configuration
func Inject() {}

// provides Dependency Injection using factory methods
func init() {
	// use the CustomLogger Implementation for our logging
	logging.New = framework.Logger
}

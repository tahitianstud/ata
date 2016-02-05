package config

import "github.com/tahitianstud/utils/logging"

// LoadConfiguration loads the dependency injection configuration
func LoadDependencies() {}

// provides Dependency Injection using factory methods
func init() {
	// use the CustomLogger Implementation for our logging
	logging.New = logging.FrameworkLogger
}

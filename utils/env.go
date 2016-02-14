package utils

import (
	s "strings"

	gotenv "github.com/subosito/gotenv"
	"os"
	"github.com/tahitianstud/ata/config"
)

var logger = config.Logger

// LoadEnvironmentFile loads the correct envfile based on the convention
// envfile name ==> .<env>
func LoadEnvironmentFile(env string) {
	envfile := config.DOT + string(os.PathSeparator) + config.DOT + s.ToLower(env)
	error := gotenv.Load(envfile)
	if error != nil {
		logger.Info("ERROR loading environment file: " + error.Error())
	}
}

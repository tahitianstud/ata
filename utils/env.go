package utils

import (
	s "strings"

	gotenv "github.com/subosito/gotenv"
)

// LoadEnvironmentFile loads the correct envfile based on the convention
// envfile name ==> .<env>
func LoadEnvironmentFile(env string) {
	envfile := s.ToLower(env)
	gotenv.Load(envfile)
}

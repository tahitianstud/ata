package utils
import (
	"github.com/tahitianstud/ata/cmd/docker"
	"github.com/tahitianstud/ata/config"
	"fmt"
	"os"
	"strings"
)

var logger = config.Logger

// AppUp determines if the app is up or not
func AppUp(app string) (bool, int) {
	numOfActiveContainers, err := docker.ActiveContainersCount(app)
	if (err != nil) {
		logger.Fatal("Unable to check if app is up or not: %s", err.Error())
	}

	if numOfActiveContainers > 0 {
		return true, numOfActiveContainers
	} else {
		return false, 0
	}
}

// PrintContainersList prints out the list of containers
func PrintContainersList(app string, idsOnly bool) {
	output, err := docker.GetContainersList(app, idsOnly)
	if err != nil {
		logger.Fatal("Unable to print out containers list %s", err.Error())
	} else {
		fmt.Println(output)
	}
}

// GuessAppFromLocation uses the location of the execution to determine the app to launch
func GuessAppFromLocation() string {
	dir, err := os.Getwd()
	if err == nil {
		separatorIndex := strings.LastIndex(dir, string(os.PathSeparator)) + 1
		if separatorIndex > -1 {
			return dir[separatorIndex:]
		}
	}

	logger.Fatal("Unable to get current working directory")
	return ""
}
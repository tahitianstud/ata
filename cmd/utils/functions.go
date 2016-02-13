package utils
import (
	"github.com/tahitianstud/ata/cmd/docker"
	"github.com/tahitianstud/ata/config"
	"fmt"
)

var logger = config.Logger

// AppUp determines if the app is up or not
func AppUp(app string) (bool, int) {
	numOfActiveContainers, err := docker.ActiveContainersCount(app)
	if (err != nil) {
		logger.Fatal("Unable to check if app is up or not: %s", err)
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

}
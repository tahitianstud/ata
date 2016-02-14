package helper
import (
	"github.com/tahitianstud/ata/cmd/docker"
	"github.com/tahitianstud/ata/config"
	"fmt"
	"strings"
	"github.com/tahitianstud/ata/utils"
	"io/ioutil"
	"os"
	"github.com/tahitianstud/utils/terminal"
	"github.com/kardianos/osext"
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

// GetAppAndEnvFromArgs checks the arguments to see if the app and env are passed
func GetAppAndEnvFromArgs(args []string) (string, string) {

	var (
		app, env string
	)

	numberOfArgs := len(args)
	logger.Trace("Found %d argument(s) to status command", numberOfArgs)

	// 3 options possibles:
	// - on a spécifié l'appli et l'environnement auquel cas on les utilise
	// - on ne précise pas l'appli mais juste l'environnement
	//   auquel cas on devine l'appli par rapport au répertoire courant
	// - on précise rien ==> environnement local
	switch numberOfArgs {
	case 2:
		app = args[0]
		env = strings.ToUpper(args[1])
	case 1:
		env = strings.ToUpper(args[0])
		app = GuessApp()
	case 0:
		env = config.DEFAULT_ENVIRONMENT
		app = GuessApp()
	}
	// TODO: validate arguments

	return app, env
}

// GuessAppFromLocation uses the location of the execution to determine the app to launch
func GuessApp() string {
	// 1) check if current directory contains .app file and load it into the ENVIRONMENT
	if _, err := ioutil.ReadFile(config.APP_FILE); err == nil {
		utils.LoadEnvironmentFile("app")
	}

	appFromFile := os.Getenv("APP")

	if len(appFromFile) > 0 {
		return appFromFile
	} else {
		// use the name of the current directory but ask for confirmation
		currentDirPath, error := osext.ExecutableFolder()
		if error != nil {
			logger.Fatal("Cannot use current directory name for the app name")
			os.Exit(1)
		} else {
			startPosition := strings.LastIndex(currentDirPath, string(os.PathSeparator)) + 1
			currentDirectoryName := currentDirPath[startPosition:]
			appNameFromDirectory := strings.ToLower(string(currentDirectoryName))

			// TODO: ask for confirmation
			fmt.Println("No app name was passed and no appfile found...")
			answer := terminal.AskForConfirmation("Do you want to use the current directory name ["+appNameFromDirectory+"] ?")
			if answer == "YES" {
				return appNameFromDirectory
			} else {
				os.Exit(0)
			}
		}

		return ""
	}
}
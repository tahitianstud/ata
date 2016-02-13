// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tahitianstud/ata/config"
	"github.com/tahitianstud/ata/cmd/utils"
	"strings"
)

var (
	logger = config.Logger
	app string
	env string
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets application status",
	Long: `Returns the application status for the <ENV> environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Trace("status called", nil)

		numberOfArgs := len(args)
		logger.Debug("Found %d argument(s) to status command\n", numberOfArgs)

		// 3 options possibles:
		// - on a spécifié l'appli et l'environnement auquel cas on les utilise
		// - on précise pas l'appli auquel cas on devine l'appli par rapport au répertoire courant
		// - on précise rien ==> environnement local

		// TODO: validate arguments

		switch numberOfArgs {
		case 2:
			app = args[0]
			env = strings.ToUpper(args[1])
		case 1:
			env = strings.ToUpper(args[0])
			app = utils.GuessAppFromLocation()
		case 0:
			app = utils.GuessAppFromLocation()
			env = DEFAULT_ENVIRONMENT
		}

		// adjust working directory
		if WorkDirectory == DEFAULT_WORK_DIRECTORY {
			WorkDirectory = DEFAULT_WORK_DIRECTORY + "/" + app
		}

		execute()
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.
}

func execute() {
	logger.Debug("Execution for trace started !")

	// TODO: implement me

	// IS APP UP ?  If yes, get the number of containers
	appIsUp, numOfContainers := utils.AppUp(app)
	if appIsUp {
		logger.Info("%s is STARTED", app, numOfContainers)
		logger.Info("There are %d containers corresponding to %s\n", numOfContainers, app)
		utils.PrintContainersList(app, false)
	} else {
		logger.Info("%s is NOT STARTED in environment %s", app, env)
	}


}

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tahitianstud/utils/logging"
	"log"
	"github.com/tahitianstud/utils/io"
	"github.com/tahitianstud/ata/config"
)

var (
	cfgFile string
	DebugMode bool
	VerboseMode bool
	WorkDirectory string
	logger = config.Logger
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ata",
	Short: "ata helps in deploying container based applications",
	Long: `ata is a tool needed in order to apply our team's convention for deploying Docker-based projects into multiple
hosts and multiple environments.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ata.yaml)")

	// add flags for debug & verbose modes
	RootCmd.PersistentFlags().BoolVarP(&DebugMode, "debug", "d", false, "activate debug mode")
	RootCmd.PersistentFlags().BoolVarP(&VerboseMode, "verbose", "v", false, "prints more output")
	RootCmd.PersistentFlags().StringVarP(&WorkDirectory, "work-directory", "w", ".", "defines work directory")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".ata")  // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// deal with verbose and debug mode
	if DebugMode {
		logger.SetLevel(logging.DebugLevel)
	}
	logger.ActivateVerboseOutput(VerboseMode)

	// check that work directory exists
	if ! io.DirectoryExists(WorkDirectory) {
		log.Fatalf("Work directory '%s' does not exist or is not readable", WorkDirectory)
	}
}

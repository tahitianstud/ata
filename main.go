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

package main

import (
	"os/exec"
	"strings"

	"github.com/tahitianstud/ata/cmd"
	"github.com/tahitianstud/ata/config"
	"github.com/tahitianstud/utils/logging"
)

const ok = "OK"

func main() {
	config.LoadDependencies()

	// TODO: check for environment conditions:

	// - detect bash
	out, bashErr := exec.Command("bash", "-c", "echo "+ok).Output()
	output := strings.TrimSpace(string(out))
	if bashErr != nil || output != ok {
		mylogger := logging.New()
		mylogger.Fatal("Bash environment not found")
	}

	// TODO: deal with embedded scripts to use

	// check opening file
	// file, fileErr := static.Open("status.sh")
	// if fileErr != nil {
	// 	log.Fatal(fileErr.Error())
	// }
	// fmt.Print(file)
	// defer file.Close()

	// if Assets == nil {
	// 	log.Die("Assets not found. Have you generated the resources ?")
	// }

	// - docker installed (check version too)
	// - docker-compose installed
	// - git installed

	cmd.Execute()
}

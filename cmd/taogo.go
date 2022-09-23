// Copyright 2022 huija
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taouniverse/tao"
	"github.com/taouniverse/taogo/cmd/project"
	"github.com/taouniverse/taogo/cmd/unit"
	"github.com/taouniverse/taogo/cmd/version"
	"log"
)

var Cmd = &cobra.Command{
	Use:   "taogo",
	Short: "Generator of tao",
	Long:  `A util to generate the universe of tao!`,
}

var data = []byte(`
tao:
  log:
    level: debug
  banner:
    hide: false
    content: 111
`)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := tao.SetAllConfigBytes(data, tao.Yaml)
	if err != nil {
		log.Fatal(err)
	}

	err = Cmd.Execute()
	if err != nil {
		tao.Fatal(err)
	}
}

func init() {
	// version
	Cmd.AddCommand(version.Cmd)
	// unit
	Cmd.AddCommand(unit.Cmd)
	// project
	Cmd.AddCommand(project.Cmd)
}

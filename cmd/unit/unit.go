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

package unit

import (
	"github.com/spf13/cobra"
	"github.com/taouniverse/taogo/constant"
	"github.com/taouniverse/taogo/tpl/git"
	"github.com/taouniverse/taogo/tpl/license"
	"github.com/taouniverse/taogo/tpl/unit"
	"github.com/taouniverse/taogo/utils"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	author  string
	module  string
	require string
	dir     string

	// Cmd of taogo unit
	Cmd = &cobra.Command{
		Use:   "unit",
		Short: "Generate unit for tao",
		Long:  `Generate unit for tao, e.g. https://github.com/taouniverse/tao-hello`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			projectName := utils.ProjectName(module)
			path, err := utils.CheckDir(dir, projectName)
			if err != nil {
				return
			}
			unitName := strings.TrimPrefix(projectName, constant.DefaultUnitPrefix)
			templates := map[string]string{
				path + ".gitignore":          git.GitIgnoreTpl,
				path + "config.go":           license.Apache2FileHeaderTpl + unit.Config,
				path + "config_test.go":      license.Apache2FileHeaderTpl + unit.ConfigTest,
				path + unitName + ".go":      license.Apache2FileHeaderTpl + unit.Unit,
				path + unitName + "_test.go": license.Apache2FileHeaderTpl + unit.UnitTest,
				path + "go.mod":              unit.Mod,
				path + "LICENSE":             license.Apache2LicenseFileTpl,
				path + "README.md":           unit.README,
			}
			params := map[string]string{
				"Author":    author,
				"Module":    module,
				"Name":      utils.PackageName(projectName),
				"Require":   require,
				"Year":      strconv.Itoa(time.Now().Year()),
				"GoVersion": runtime.Version(),
			}
			err = utils.ExecuteTemplate(templates, params)
			if err != nil {
				return
			}
			err = utils.ModTidy(path)
			if err != nil {
				return
			}
			return utils.TestCover(path)
		},
	}
)

func init() {
	// Persistence Flags
	Cmd.PersistentFlags().StringVarP(&module, "module", "m", "github.com/taouniverse/tao-hello", "target module name of unit")
	Cmd.PersistentFlags().StringVarP(&require, "require", "r", "", "require modules, split by '"+constant.ParamSplit+"'")
	Cmd.PersistentFlags().StringVarP(&dir, "dir", "d", "./", "unit's parent path")
	Cmd.PersistentFlags().StringVarP(&author, "author", "a", "huija", "author of the target unit")
}

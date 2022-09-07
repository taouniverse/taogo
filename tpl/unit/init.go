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

// Init init.go
const Init = `
package {{ .Module | package }}

import (
	"encoding/json"
	"github.com/taouniverse/tao"
	// TODO use the following modules to develop this unit
	{{ .Require | import }}
)

/**
import _ "{{ .Module }}"
*/

// {{ .Module | package | first | upper }} config of {{ .Module | package }}
var {{ .Module | package | first | upper }} = new(Config)

func init() {
	err := tao.Register(ConfigKey, func() (err error) {
		// 1. transfer config bytes to object
		bytes, err := tao.GetConfigBytes(ConfigKey)
		if err != nil {
			{{ .Module | package | first | upper }} = {{ .Module | package | first | upper }}.Default().(*Config)
		} else {
			err = json.Unmarshal(bytes, &{{ .Module | package | first | upper }})
			if err != nil {
				return err
			}
		}
		// 2. set object to tao
		err = tao.SetConfig(ConfigKey, {{ .Module | package | first | upper }})
		if err != nil {
			return err
		}
		// TODO setup your unit before run JOB
		return
	})
	if err != nil {
		panic(err.Error())
	}
}`

// Copyright (c) 2022 Trim21 <trim21.me@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

// A help script to generate json schema from go struct
package main

import (
	"fmt"
	"reflect"

	"github.com/danielgtaylor/huma/schema"
	"github.com/goccy/go-json"
	"gopkg.in/yaml.v3"

	"github.com/bangumi/server/web/req"
)

func main() {
	// Then later you can do:
	s, err := schema.Generate(reflect.TypeOf(req.PutSubject{}))
	if err != nil {
		panic(err)
	}

	fmt.Println(dump(s))
}

func dump(v interface{}) string {
	jsonRaw, err := json.MarshalNoEscape(v)
	if err != nil {
		panic(err)
	}

	var d interface{}
	if err = json.Unmarshal(jsonRaw, &d); err != nil {
		panic(err)
	}

	yamlRaw, err := yaml.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(yamlRaw)
}

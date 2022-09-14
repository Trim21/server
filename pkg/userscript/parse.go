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

package userscript

import (
	"fmt"
	"regexp"

	"github.com/bangumi/server/pkg/userscript/internal/match"
)

type UserScript struct {
	Meta    Meta
	Content string // js script content
}

var pattern = regexp.MustCompile(
	`\B// ==UserScript==\r?\n(?P<meta>[\S\s]*?)\r?\n// ==/UserScript==(?P<content>[\S\s]*)$`,
)

func Parse(s string) (UserScript, error) {
	groups := pattern.FindStringSubmatch(s)

	metaRaw := groups[1]
	content := groups[2]

	meta, err := parseMeta(metaRaw)
	if err != nil {
		return UserScript{}, err
	}

	return UserScript{
		Meta:    meta,
		Content: content,
	}, nil
}

const KeyMatch = "match"
const KeyInclude = "include"

func (u UserScript) Match(url string) bool {
	includes := u.Meta.GetAll(KeyInclude)
	for _, include := range includes {
		if include[0] == '/' && include[len(include)-1] == '/' {
			ok, _ := regexp.MatchString(include, url)
			if ok {
				fmt.Println(include)
				return true
			}
		}
	}

	for _, matchRule := range u.Meta.GetAll(KeyMatch) {
		if match.Match(url, matchRule) {
			return true
		}
	}

	return false
}

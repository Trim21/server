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
	"net/url"
	"regexp"
	"strings"
)

var metaPattern = regexp.MustCompile(`@([\w-]+)\s+(.+)`)

func parseMeta(s string) (Meta, error) {
	m := Meta{u: make(url.Values, strings.Count(s, "\n"))}

	lines := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
	for _, line := range lines {
		meta := metaPattern.FindStringSubmatch(line)
		m.Add(meta[1], meta[2])
	}

	return m, nil
}

type Meta struct {
	u url.Values
}

func (m Meta) Get(key string) string {
	return m.u.Get(key)
}

func (m Meta) Set(key, value string) {
	m.u.Set(key, value)
}

func (m Meta) Add(key, value string) {
	m.u.Add(key, value)
}

func (m Meta) Del(key string) {
	m.u.Del(key)
}

func (m Meta) Has(key string) bool {
	return m.u.Has(key)
}

func (m Meta) GetAll(key string) []string {
	return m.u[key]
}

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

package main

import (
	"fmt"

	"github.com/bangumi/server/pkg/userscript"
)

const s = `// ==UserScript==
// @name        Awesome Script
// @description This script even does the laundry!
// @downloadURL https://www.example.com/myscript.user.js
// @homepageURL https://github.com/gantt/downloadyoutube
// @author      Gantt
// @version     1.8.3
// @date        2015-05-17
// @include     https://www.youtube.com/*
// @exclude     https://www.youtube.com/embed/*
// @match       https://www.youtube.com/*
// @grant       GM_xmlhttpRequest
// @grant       GM_getValue
// @grant       GM_setValue
// @run-at      document-end
// @license     MIT License
// ==/UserScript==

var whoami = "USERSCRIPT"
`

func main() {
	u, err := userscript.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Match("https://www.youtube.com"))
}

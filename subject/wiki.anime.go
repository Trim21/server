// Copyright (c) 2022 Sociosarbis <136657577@qq.com>
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

package subject

import (
	"github.com/bangumi/server/internal/strparse"
	"github.com/bangumi/server/pkg/wiki"
)

func extractAnimeWiki(w wiki.Wiki) extractedWikiData {
	var e extractedWikiData
	for _, field := range w.Fields {

		if field.Null {
			continue
		}

		switch field.Key {
		case cnNameKey:
			e.NameCN = field.Value

		case keyStart, "放送开始":
			e.Date = extractDateString(field)

		case "放送星期":
			e.Airtime = parseAirtime(field.Value)

		case keyEps:
			if eps, err := strparse.Uint32(field.Value); err != nil {
				e.Eps = eps
			}
		}

	}

	return e
}

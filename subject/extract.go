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
	"regexp"
	"strconv"
	"time"

	"github.com/bangumi/server/config"
	"github.com/bangumi/server/model"
	"github.com/bangumi/server/pkg/wiki"
)

type extractedWikiData struct {
	Date    time.Time
	NameCN  string
	Airtime uint8
}

func extractFromWiki(t model.SubjectType, w wiki.Wiki) extractedWikiData {
	switch t {
	case model.SubjectAnime:
		return extractAnimeWiki(w)
	case model.SubjectReal:
		return extractRealWiki(w)
	case model.SubjectMusic:
		return extractMusicWiki(w)
	case model.SubjectGame:
		return extractGameWiki(w)
	case model.SubjectBook:
		return extractBookWiki(w)
	}

	return extractedWikiData{}
}

const cnNameKey = "中文名"
const keyStart = "开始"

var datePattern = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})`)

func extractDateString(f wiki.Field) time.Time {
	var raw = f.Value
	if f.Array {
		raw = f.Values[0].Value
	}

	p := datePattern.FindStringSubmatch(raw)
	if len(p) == 0 {
		return time.Time{}
	}

	year, err := strconv.Atoi(p[1])
	if err != nil {
		return time.Time{}
	}

	month, err := strconv.Atoi(p[2])
	if err != nil {
		return time.Time{}
	}
	if month == 0 {
		month = 1
	}

	day, err := strconv.Atoi(p[3])
	if err != nil {
		return time.Time{}
	}

	if day == 0 {
		day = 1
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, config.TZ)
}

//nolint:gomnd
func parseAirtime(s string) uint8 {
	switch s {
	case "1", "星期一", "周一", "一":
		return 1

	case "2", "星期二", "周二", "二":
		return 2

	case "3", "星期三", "周三", "三":
		return 3

	case "4", "星期四", "周四", "四":
		return 4

	case "5", "星期五", "周五", "五":
		return 5

	case "6", "星期六", "周六", "六":
		return 6

	case "7", "星期日", "周日", "日":
		return 7
	}

	return 0
}

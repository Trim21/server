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

package req

type PutSubject struct {
	Name        string `json:"name"`
	Infobox     string `json:"infobox"`
	Summary     string `json:"summary" doc:"条目简介"`
	EditSummary string `json:"edit_summary" doc:"编辑摘要"`
	Entry       bool   `json:"entry" doc:"if this subject is main entry for a series"`
	Platform    uint16 `json:"platform"`
	NSFW        bool   `json:"nsfw"`
}

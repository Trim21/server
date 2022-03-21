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

package model_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/bangumi/server/config"
	"github.com/bangumi/server/model"
)

func TestSubject_DateString(t *testing.T) {
	t.Parallel()

	assert.Equal(t, (*string)(nil), model.Subject{}.DateString())

	s := model.Subject{
		Date: time.Date(2005, 9, 1, 0, 0, 0, 0, config.TZ),
	}

	assert.Equal(t, "2005-09-01", *(s.DateString()))
}

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

package handler_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/internal/test"
	"github.com/bangumi/server/mocks"
	"github.com/bangumi/server/model"
	"github.com/bangumi/server/web/req"
	"github.com/bangumi/server/web/res"
)

type mockAuth struct{ u domain.Auth }

func (m mockAuth) GetByToken(ctx context.Context, token string) (domain.Auth, error) {
	return m.u, nil
}

func TestHappyPath(t *testing.T) {
	t.Parallel()
	m := &mocks.SubjectRepo{}
	m.EXPECT().Get(mock.Anything, uint32(7)).Return(model.Subject{ID: 7}, nil)

	app := test.GetWebApp(t,
		test.Mock{
			AuthRepo:    mockAuth{domain.Auth{RegTime: time.Unix(1e10, 0)}},
			SubjectRepo: m,
		},
	)

	req := httptest.NewRequest(http.MethodGet, "/v0/subjects/7", http.NoBody)
	req.Header.Set("authorization", "Bearer token")

	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var r res.SubjectV0
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&r))
	require.Equal(t, uint32(7), r.ID)
}

func TestNSFW_200(t *testing.T) {
	t.Parallel()

	m := &mocks.SubjectRepo{}
	m.EXPECT().Get(mock.Anything, uint32(7)).Return(model.Subject{NSFW: true}, nil)

	app := test.GetWebApp(t,
		test.Mock{
			AuthRepo:    mockAuth{domain.Auth{ID: 1, RegTime: time.Unix(1e9, 0)}},
			SubjectRepo: m,
		},
	)

	req := httptest.NewRequest(http.MethodGet, "/v0/subjects/7", http.NoBody)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer token")

	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "200 for authorized user")
}

func TestNSFW_404(t *testing.T) {
	t.Parallel()

	m := &mocks.SubjectRepo{}
	m.EXPECT().Get(mock.Anything, uint32(7)).Return(model.Subject{NSFW: true}, nil)

	app := test.GetWebApp(t,
		test.Mock{
			AuthRepo:    mockAuth{domain.Auth{}},
			SubjectRepo: m,
		},
	)

	req := httptest.NewRequest(http.MethodGet, "/v0/subjects/7", http.NoBody)
	req.Header.Set("authorization", "Bearer token")

	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusNotFound, resp.StatusCode, "404 for unauthorized user")
}

func Test_web_subject_Redirect(t *testing.T) {
	t.Parallel()
	m := &mocks.SubjectRepo{}
	m.EXPECT().Get(mock.Anything, uint32(8)).Return(model.Subject{Redirect: 2}, nil)

	app := test.GetWebApp(t,
		test.Mock{
			SubjectRepo: m,
		},
	)

	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/v0/subjects/8", http.NoBody))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusFound, resp.StatusCode, "302 for redirect repository")
	require.Equal(t, "/v0/subjects/2", resp.Header.Get("location"))
}

func Test_web_subject_bad_id(t *testing.T) {
	t.Parallel()
	m := &mocks.SubjectRepo{}

	app := test.GetWebApp(t, test.Mock{SubjectRepo: m})

	for _, path := range []string{"/v0/subjects/0", "/v0/subjects/-1", "/v0/subjects/a"} {
		path := path
		t.Run(path, func(t *testing.T) {
			t.Parallel()
			resp, err := app.Test(httptest.NewRequest(http.MethodGet, path, http.NoBody))
			require.NoError(t, err)
			defer resp.Body.Close()

			require.Equal(t, http.StatusBadRequest, resp.StatusCode, "400 for redirect subject id")
		})
	}
}

func TestHandler_PutSubject(t *testing.T) {
	t.Parallel()
	m := &mocks.SubjectRepo{}
	defer m.AssertExpectations(t)

	user := &mocks.AuthRepo{}
	user.EXPECT().GetByToken(mock.Anything, mock.Anything).Return(domain.Auth{
		RegTime: time.Time{},
		ID:      1,
		Group:   1,
	}, nil)

	m.EXPECT().Set(mock.Anything, uint32(363612), mock.Anything, mock.Anything).Return(nil)
	m.EXPECT().Get(mock.Anything, uint32(363612)).Return(model.Subject{}, nil)
	app := test.GetWebApp(t, test.Mock{SubjectRepo: m, AuthRepo: user})

	var b = bytes.NewBuffer(nil)
	err := json.NewEncoder(b).Encode(req.PutSubject{
		Name: "n",
		Infobox: `
{{Infobox
|a=1
|开始=2020-10-03
}}`,
		Summary:     "summary",
		EditSummary: "commit message",
		Platform:    2,
	})
	require.NoError(t, err)

	request := httptest.NewRequest(http.MethodPut, "/v0/subjects/363612", b)
	request.Header.Set(fiber.HeaderAuthorization, "Bearer 1")

	resp, err := app.Test(request)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, http.StatusNoContent, resp.StatusCode, string(body))
}

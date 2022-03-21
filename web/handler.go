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

package web

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/uber-go/tally/v4"

	"github.com/bangumi/server/web/handler"
	"github.com/bangumi/server/web/res"
	"github.com/bangumi/server/web/util"
)

// ResistRouter add all router and default 404 Handler to app.
func ResistRouter(app *fiber.App, h handler.Handler, scope tally.Scope) {
	app.Use(handler.MiddlewareAccessUser(h))

	app.Get("/v0/subjects/:id", h.GetSubject)
	app.Put("/v0/subjects/:id", h.PutSubject)
	app.Get("/v0/subjects/:id/persons", h.GetSubjectRelatedPersons)
	app.Get("/v0/subjects/:id/subjects", h.GetSubjectRelatedSubjects)
	app.Get("/v0/subjects/:id/characters", h.GetSubjectRelatedCharacters)
	app.Get("/v0/persons/:id", h.GetPerson)
	app.Get("/v0/persons/:id/subjects", h.GetPersonRelatedSubjects)
	app.Get("/v0/persons/:id/characters", h.GetPersonRelatedCharacters)
	app.Get("/v0/characters/:id", h.GetCharacter)
	app.Get("/v0/characters/:id/subjects", h.GetCharacterRelatedSubjects)
	app.Get("/v0/characters/:id/persons", h.GetCharacterRelatedPersons)
	app.Get("/v0/episodes/:id", h.GetEpisode)
	app.Get("/v0/episodes", h.ListEpisode)
	app.Get("/v0/me", h.GetCurrentUser)
	app.Get("/v0/users/:username/collections", h.ListCollection)
	app.Get("/v0/indices/:id", h.GetIndex)
	app.Get("/v0/indices/:id/subjects", h.GetIndexSubjects)

	app.Get("/v0/revisions/persons/:id", h.GetPersonRevision)
	app.Get("/v0/revisions/persons", h.ListPersonRevision)

	// 给所有的 path 添加 metrics.
	for _, routes := range app.Stack() {
		for i, r := range routes {
			realHandler := r.Handlers[len(r.Handlers)-1]
			if !strings.HasPrefix(utils.FunctionName(realHandler), "github.com/bangumi/server/web/handler.Handler.") {
				continue
			}

			handlers := make([]fiber.Handler, len(r.Handlers)+1)
			for j := 0; j < len(r.Handlers)-1; j++ {
				handlers[j] = r.Handlers[j]
			}

			reqCounter := scope.Tagged(map[string]string{"path": r.Path, "method": r.Method}).Counter("request_count")
			handlers[len(handlers)-2] = func(c *fiber.Ctx) error {
				reqCounter.Inc(1)
				return c.Next()
			}

			handlers[len(handlers)-1] = realHandler
			routes[i].Handlers = handlers
		}
	}

	// default 404 Handler, all router should be added before this router
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(res.Error{
			Title:       "Not Found",
			Description: "This is default response, if you see this response, please check your request path",
			Details:     util.DetailFromRequest(c),
		})
	})
}

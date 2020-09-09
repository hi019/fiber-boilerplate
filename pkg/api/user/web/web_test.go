package web_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hi019/fiber-boilerplate/ent"
	"github.com/hi019/fiber-boilerplate/ent/enttest"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog"

	_ "github.com/mattn/go-sqlite3"

	us "github.com/hi019/fiber-boilerplate/pkg/api/user"
	"github.com/hi019/fiber-boilerplate/pkg/api/user/web"
)

type user struct {
	Email string
	ID    int
}

func createAPI(t *testing.T) (*fiber.App, *ent.Client) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	app := fiber.New()
	logger := &zerolog.Logger{}
	svc := us.Initialize(client, logger)
	vd := validator.New()

	web.NewHttp(svc, app, vd)

	return app, client
}

func TestSignup(t *testing.T) {
	app, db := createAPI(t)
	defer db.Close()

	cases := []struct {
		name       string
		req        string
		wantStatus int
		wantResp   *user
	}{
		{
			name:       "Create a user",
			req:        `{"email": "test@email.com", "password": "password"}`,
			wantStatus: fiber.StatusOK,
		},
		{
			name:       "Fail creating a user with invalid email",
			req:        `{"email": "test", "password": "password"}`,
			wantStatus: fiber.StatusBadRequest,
		},
		{
			name:       "Fail creating a user that already exists",
			req:        `{"email": "test@email.com", "password": "password"}`,
			wantStatus: fiber.StatusBadRequest,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/signup", strings.NewReader(tt.req))
			req.Header.Add("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if tt.wantResp != nil {
				response := new(user)
				if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.wantResp, response)
			}

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}


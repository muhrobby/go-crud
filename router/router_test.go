package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestShowUser(t *testing.T) {

	app := fiber.New()

	Routes(app)

	t.Run("GET /api/users", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
		resp, err := app.Test(req, -1)

		assert.Nil(t, err, "error getting")
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Seharusnyaa 200")

		body, err := io.ReadAll(resp.Body)

		assert.Nil(t, err, "error reading")
		assert.Equal(t, "{\"message\":\"Hello User\"}", string(body), "ok")
	})

}

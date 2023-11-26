package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

func TestShowUser(t *testing.T) {

	app.Get("/users", ShowUser)

	req := httptest.NewRequest("GET", "/users", nil)

	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	errR := json.NewDecoder(resp.Body).Decode(&result)

	assert.Nil(t, errR)

	assert.Equal(t, "success", result["status"])
	assert.Equal(t, "berhasil ambil data", result["message"])

}

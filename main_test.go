package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mhdiiilham/ginorm/routers"
	"github.com/stretchr/testify/assert"
	log "github.com/sirupsen/logrus"
)

func TestTodoWithoutAuthentication(t *testing.T) {
	var res map[string]string

	r      := routers.Router()
	w      := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos", nil)
	r.ServeHTTP(w, req)
	json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "NOT AUTHORIZED", res["errors"])
}

func TestTodoWithToken(t *testing.T) {
	var res map[string]string

	r      := routers.Router()
	w      := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todos", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlLCJleHAiOjE1OTI2NjIwNzAsInVzZXJfZW1haWwiOiJzdXBlcmFkbWluQG11aGFtbWFkaWxoYW0uY29tIiwidXNlcl9pZCI6NH0.og0_9sCDUZfLRNGgwcqSqZEFIwGwEWtZpPaXvDUwZ-U")
	r.ServeHTTP(w, req)
	json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, http.StatusOK, w.Code)
}
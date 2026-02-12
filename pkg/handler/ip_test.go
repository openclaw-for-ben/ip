// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bborbe/ip/pkg/handler"
)

func TestIPHandler_XForwardedFor(t *testing.T) {
	h := handler.NewIPHandler()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.195, 70.41.3.18, 150.172.238.178")

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if rec.Body.String() != "203.0.113.195" {
		t.Errorf("expected '203.0.113.195', got '%s'", rec.Body.String())
	}
}

func TestIPHandler_XRealIP(t *testing.T) {
	h := handler.NewIPHandler()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Real-IP", "192.168.1.1")

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if rec.Body.String() != "192.168.1.1" {
		t.Errorf("expected '192.168.1.1', got '%s'", rec.Body.String())
	}
}

func TestIPHandler_RemoteAddr(t *testing.T) {
	h := handler.NewIPHandler()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// httptest.NewRequest sets RemoteAddr to "192.0.2.1:1234"

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if rec.Body.String() != "192.0.2.1" {
		t.Errorf("expected '192.0.2.1', got '%s'", rec.Body.String())
	}
}

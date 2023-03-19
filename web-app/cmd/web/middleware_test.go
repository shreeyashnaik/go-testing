package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	var app application

	// create a dummy handler, to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure context has some value
		val := r.Context().Value(ctxUserKey)
		if val == nil {
			t.Error(ctxUserKey, "not present")
		}

		// make sure context value is a string
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}

		t.Log(ip)
	})

	for _, e := range tests {
		// create a handler to test
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)
		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	tests := []struct {
		name          string
		ctx           context.Context
		expectedIPstr string
	}{
		{"valid", context.WithValue(context.Background(), ctxUserKey, "192.0.2.1"), "192.0.2.1"},
	}

	// create an app variable
	app := application{}

	for _, e := range tests {
		if res := app.ipFromContext(e.ctx); res != e.expectedIPstr {
			t.Errorf("for %s wrong value returned from ctx: expected %s, got %s", e.name, e.expectedIPstr, res)
		}
	}
}

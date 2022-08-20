package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ats2otus/final_project/pkg/bwlist"
)

func Test_rateService_append_black_white_list(t *testing.T) {
	testCases := []struct {
		name string
		body string
		code int
	}{
		{
			name: "bad json payload",
			body: "]",
			code: http.StatusBadRequest,
		},
		{
			name: "bad subnet",
			body: "{\"subnet\":\"h.e.l.l.o\"}",
			code: http.StatusBadRequest,
		},
		{
			name: "ip instead of subnet",
			body: "{\"subnet\":\"127.0.0.1\"}",
			code: http.StatusBadRequest,
		},
		{
			name: "correct case",
			body: "{\"subnet\":\"192.168.0.5/24\"}",
			code: http.StatusAccepted,
		},
	}

	rs := rateService{
		blacklist: bwlist.New(),
		whitelist: bwlist.New(),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			rs.appendBlacklist(rw, req)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			rs.appendWhitelist(rw, req)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}
		})
	}
}

func Test_rateService_remove_black_white_list(t *testing.T) {
	testCases := []struct {
		name   string
		subnet string
		code   int
	}{
		{
			name:   "empty subnet",
			subnet: "",
			code:   http.StatusBadRequest,
		},
		{
			name:   "wrong subnet",
			subnet: "h.e.l.l.o",
			code:   http.StatusBadRequest,
		},
		{
			name:   "ip instead of subnet",
			subnet: "127.0.0.1",
			code:   http.StatusBadRequest,
		},
		{
			name:   "correct case",
			subnet: "192.168.0.5/24",
			code:   http.StatusAccepted,
		},
	}

	rs := rateService{
		blacklist: bwlist.New(),
		whitelist: bwlist.New(),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/?subnet="+tc.subnet, http.NoBody)
			if err != nil {
				t.Fatal(err)
			}

			rs.removeBlacklist(rw, req)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/?subnet="+tc.subnet, http.NoBody)
			if err != nil {
				t.Fatal(err)
			}

			rs.removeWhitelist(rw, req)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}
		})
	}
}

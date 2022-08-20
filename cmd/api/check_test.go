package main

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_rateService_allow(t *testing.T) {
	testCases := []struct {
		name    string
		body    string
		code    int
		answer  bool
		service rateService
	}{
		{
			name: "bad json payload",
			body: "]",
			code: http.StatusBadRequest,
		},
		{
			name:   "ip in whitelist",
			body:   `{"ip":"0.0.0.0"}`,
			code:   http.StatusOK,
			answer: true,
			service: rateService{
				whitelist: mockBWList{true},
				blacklist: mockBWList{false},
			},
		},
		{
			name:   "ip in blacklist",
			body:   `{"ip":"0.0.0.0"}`,
			code:   http.StatusOK,
			answer: false,
			service: rateService{
				whitelist: mockBWList{false},
				blacklist: mockBWList{true},
			},
		},
		{
			name:   "ip limit reached",
			body:   `{"ip":"0.0.0.0"}`,
			code:   http.StatusOK,
			answer: false,
			service: rateService{
				whitelist:     mockBWList{false},
				blacklist:     mockBWList{false},
				limitByIP:     mockLimiter{false},
				limitByLogin:  mockLimiter{true},
				limitByPasswd: mockLimiter{true},
			},
		},
		{
			name:   "login limit reached",
			body:   `{"ip":"0.0.0.0"}`,
			code:   http.StatusOK,
			answer: false,
			service: rateService{
				whitelist:     mockBWList{false},
				blacklist:     mockBWList{false},
				limitByIP:     mockLimiter{true},
				limitByLogin:  mockLimiter{false},
				limitByPasswd: mockLimiter{true},
			},
		},
		{
			name:   "password limit reached",
			body:   `{"ip":"0.0.0.0"}`,
			code:   http.StatusOK,
			answer: false,
			service: rateService{
				whitelist:     mockBWList{false},
				blacklist:     mockBWList{false},
				limitByIP:     mockLimiter{true},
				limitByLogin:  mockLimiter{true},
				limitByPasswd: mockLimiter{false},
			},
		},
		{
			name:   "correct case",
			body:   `{"ip":"0.0.0.0"}`,
			code:   http.StatusOK,
			answer: true,
			service: rateService{
				whitelist:     mockBWList{false},
				blacklist:     mockBWList{false},
				limitByIP:     mockLimiter{true},
				limitByLogin:  mockLimiter{true},
				limitByPasswd: mockLimiter{true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			tc.service.allow(rw, req)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}

			var result Result
			if err := json.NewDecoder(rw.Body).Decode(&result); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}

			if result.Ok != tc.answer {
				t.Fatalf("unexpected answer; expected: %t actual: %t", tc.answer, result.Ok)
			}
		})
	}
}

func Test_rateService_reset(t *testing.T) {
	testCases := []struct {
		name    string
		body    string
		code    int
		answer  bool
		service rateService
	}{
		{
			name: "bad json payload",
			body: "]",
			code: http.StatusBadRequest,
		},
		{
			name: "correct case",
			body: `{"ip":"0.0.0.0","login":"freak192"}`,
			code: http.StatusAccepted,
		},
	}

	rs := rateService{
		limitByIP:    mockLimiter{false},
		limitByLogin: mockLimiter{true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			rs.reset(rw, req)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}
		})
	}
}

type mockBWList struct{ result bool }

func (m mockBWList) Contains(net.IP) bool { return m.result }
func (mockBWList) Append(*net.IPNet)      {}
func (mockBWList) Remove(*net.IPNet)      {}

type mockLimiter struct{ result bool }

func (m mockLimiter) Allow(key string) bool { return m.result }
func (m mockLimiter) Reset(key string)      {}

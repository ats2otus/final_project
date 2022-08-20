package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_rateService_writeError(t *testing.T) {
	var service rateService

	testCases := []struct {
		name  string
		code  int
		err   error
		empty bool
	}{
		{
			name:  "204 no content",
			code:  http.StatusNoContent,
			empty: true,
			err:   http.ErrAbortHandler,
		},
		{
			name:  "error is nil",
			code:  http.StatusInternalServerError,
			empty: true,
			err:   nil,
		},
		{
			name:  "result with body",
			code:  http.StatusInternalServerError,
			empty: false,
			err:   http.ErrAbortHandler,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
			if err != nil {
				t.Fatal(err)
			}

			service.writeError(rw, req, tc.code, tc.err)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}
			if tc.empty && rw.Body.Len() != 0 {
				t.Fatalf("unexpected res.body; expected: empty actual: %s", rw.Body.String())
			}
		})
	}
}

func Test_rateService_writeResult(t *testing.T) {
	var service rateService

	testCases := []struct {
		name  string
		code  int
		body  interface{}
		empty bool
	}{
		{
			name:  "204 no content",
			code:  http.StatusNoContent,
			empty: true,
			body:  struct{}{},
		},
		{
			name:  "body is nil",
			code:  http.StatusOK,
			empty: true,
			body:  nil,
		},
		{
			name:  "result with body",
			code:  http.StatusOK,
			empty: false,
			body:  struct{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
			if err != nil {
				t.Fatal(err)
			}

			service.writeResult(rw, req, tc.code, tc.body)

			if rw.Code != tc.code {
				t.Fatalf("unexpected res.code; expected: %d actual: %d", tc.code, rw.Code)
			}
			if tc.empty && rw.Body.Len() != 0 {
				t.Fatalf("unexpected res.body; expected: empty actual: %s", rw.Body.String())
			}
		})
	}
}

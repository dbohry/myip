package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientIP(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		headers    map[string]string
		want       string
	}{
		{"remote addr ipv4", "1.2.3.4:5678", nil, "1.2.3.4"},
		{"remote addr ipv6", "[::1]:5678", nil, "::1"},
		{"x-forwarded-for single", "1.2.3.4:5678", map[string]string{"X-Forwarded-For": "5.6.7.8"}, "5.6.7.8"},
		{"x-forwarded-for multiple", "1.2.3.4:5678", map[string]string{"X-Forwarded-For": "5.6.7.8, 9.10.11.12"}, "5.6.7.8"},
		{"x-real-ip", "1.2.3.4:5678", map[string]string{"X-Real-IP": "9.9.9.9"}, "9.9.9.9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.RemoteAddr = tt.remoteAddr
			for k, v := range tt.headers {
				r.Header.Set(k, v)
			}
			if got := clientIP(r); got != tt.want {
				t.Errorf("clientIP() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGetPublicIP(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.RemoteAddr = "1.2.3.4:5678"
	w := httptest.NewRecorder()

	getPublicIP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}
	want := `{"ip":"1.2.3.4"}`
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

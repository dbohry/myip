package main

import "testing"

func TestParseUserAgent(t *testing.T) {
	tests := []struct {
		name        string
		ua          string
		wantBrowser string
		wantOS      string
	}{
		{
			name:        "chrome windows",
			ua:          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			wantBrowser: "Chrome 120.0.0.0",
			wantOS:      "Windows",
		},
		{
			name:        "firefox linux",
			ua:          "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0",
			wantBrowser: "Firefox 121.0",
			wantOS:      "Linux",
		},
		{
			name:        "safari macos",
			ua:          "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
			wantBrowser: "Safari 17.0",
			wantOS:      "macOS",
		},
		{
			name:        "safari ios",
			ua:          "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
			wantBrowser: "Safari 17.0",
			wantOS:      "iOS",
		},
		{
			name:        "chrome android",
			ua:          "Mozilla/5.0 (Linux; Android 14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			wantBrowser: "Chrome 120.0.0.0",
			wantOS:      "Android",
		},
		{
			name:        "edge windows",
			ua:          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
			wantBrowser: "Edge 120.0.0.0",
			wantOS:      "Windows",
		},
		{
			name:        "empty",
			ua:          "",
			wantBrowser: "Unknown",
			wantOS:      "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser, os := parseUserAgent(tt.ua)
			if browser != tt.wantBrowser {
				t.Errorf("browser = %q, want %q", browser, tt.wantBrowser)
			}
			if os != tt.wantOS {
				t.Errorf("os = %q, want %q", os, tt.wantOS)
			}
		})
	}
}

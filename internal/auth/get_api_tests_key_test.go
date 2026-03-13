package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		wantKey     string
		wantErr     bool
	}{
		{
			name:        "missing authorization header",
			headerValue: "",
			wantKey:     "ApiKey abc123",
			wantErr:     true,
		},
		{
			name:        "malformed authorization header",
			headerValue: "badkey",
			wantKey:     "ApiKey abc123",
			wantErr:     true,
		},
		{
			name:        "valid api key header",
			headerValue: "ApiKey abc123",
			wantKey:     "ApiKey abc123",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.headerValue != "" {
				headers.Set("Authorization", tt.headerValue)
			}

			gotKey, err := GetAPIKey(headers)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("did not expect error, got: %v", err)
			}

			if gotKey != tt.wantKey {
				t.Fatalf("expected key %q, got %q", tt.wantKey, gotKey)
			}
		})
	}
}

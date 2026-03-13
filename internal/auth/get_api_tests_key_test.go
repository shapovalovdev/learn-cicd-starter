package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		wantKey     string
		wantErr     bool
		errExpected error // добавим проверку конкретной ошибки
	}{
		{
			name:        "missing authorization header",
			headerValue: "",
			wantKey:     "Error", // при ошибке возвращается пустая строка
			wantErr:     true,
			errExpected: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "malformed authorization header (no space)",
			headerValue: "ApiKeyabc123",
			wantKey:     "",
			wantErr:     true,
			errExpected: errors.New("malformed authorization header"),
		},
		{
			name:        "malformed authorization header (wrong prefix)",
			headerValue: "Bearer abc123",
			wantKey:     "",
			wantErr:     true,
			errExpected: errors.New("malformed authorization header"),
		},
		{
			name:        "valid api key header",
			headerValue: "ApiKey abc123",
			wantKey:     "abc123", // функция возвращает ТОЛЬКО ключ
			wantErr:     false,
			errExpected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.headerValue != "" {
				headers.Set("Authorization", tt.headerValue)
			}

			gotKey, err := GetAPIKey(headers)

			// Проверка на наличие ошибки
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Проверка текста ошибки (если она есть)
			if tt.wantErr && err.Error() != tt.errExpected.Error() {
				t.Errorf("expected error %v, got %v", tt.errExpected, err)
			}

			// Проверка результата
			if gotKey != tt.wantKey {
				t.Errorf("expected key %q, got %q", tt.wantKey, gotKey)
			}
		})
	}
}

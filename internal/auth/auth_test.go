package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name        string
		headerVal   string
		expectedKey string
		expectedErr bool
	}{
		{
			name:        "valid api key",
			headerVal:   "ApiKey valid-api-key",
			expectedKey: "valid-api-key",
			expectedErr: false,
		},
		{
			name:        "no auth header",
			headerVal:   "",
			expectedKey: "",
			expectedErr: true,
		},
		{
			name:        "malformed header",
			headerVal:   "BadPrefix some-key",
			expectedKey: "",
			expectedErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			if tc.headerVal != "" {
				req.Header.Set("Authorization", tc.headerVal)

			}
			key, err := GetAPIKey(req.Header)
			if tc.expectedErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if key != tc.expectedKey {
				t.Fatalf("expected key %q but got %q", tc.expectedKey, key)
			}
		})
	}
}

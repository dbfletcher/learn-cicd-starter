package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Define a series of test cases
	cases := []struct {
		name          string
		header        http.Header
		expectedKey   string
		expectError   bool
		expectedError error
	}{
		{
			name: "Valid API Key",
			header: http.Header{
				"Authorization": []string{"ApiKey my-secret-api-key"},
			},
			expectedKey: "my-secret-api-key",
			expectError: false,
		},
		{
			name:          "No Authorization Header",
			header:        http.Header{},
			expectedKey:   "",
			expectError:   true,
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header - Wrong Scheme",
			header: http.Header{
				"Authorization": []string{"Bearer my-secret-api-key"},
			},
			expectedKey: "",
			expectError: true,
		},
		/*
			{
				name: "Malformed Authorization Header - No Key",
				header: http.Header{
					"Authorization": []string{"ApiKey "},
				},
				expectedKey: "",
				expectError: true,
			},
		*/
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actualKey, err := GetAPIKey(tc.header)

			if actualKey != tc.expectedKey {
				t.Errorf("expected key '%s', got '%s'", tc.expectedKey, actualKey)
			}

			if tc.expectError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
				if tc.expectedError != nil && err != tc.expectedError {
					t.Errorf("expected error '%v', got '%v'", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got one: %v", err)
				}
			}
		})
	}
}


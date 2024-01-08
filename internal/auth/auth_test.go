package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case 1: No Authorization header
	headers := make(http.Header)
	apiKey, err := GetAPIKey(headers)
	if apiKey != "" || err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v and %v", apiKey, err)
	}

	// Test case 2: Malformed Authorization header
	headers.Set("Authorization", "Bearer123")
	apiKey, err = GetAPIKey(headers)
	if apiKey != "" || err == nil {
		t.Errorf("Expected error due to malformed header, got %v and %v", apiKey, err)
	}

	// Test case 3: Correct Authorization header
	headers.Set("Authorization", "ApiKey my-api-key")
	apiKey, err = GetAPIKey(headers)
	if apiKey != "my-api-key" || err != nil {
		t.Errorf("Expected API key 'my-api-key', got %v and %v", apiKey, err)
	}
}

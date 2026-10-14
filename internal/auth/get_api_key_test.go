package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("should return a valid API key", func(t *testing.T) {
		headers := http.Header{
			"Authorization": []string{"ApiKey your_expected_api_key"},
		}
		expected := "your_expected_api_key"
		actual, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})
}

func TestEmptyHeaders(t *testing.T) {
	t.Run("should return an error for empty headers", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
	})
}

func TestMalformedAuthorizationHeader(t *testing.T) {
	t.Run("should return an error for malformed authorization header", func(t *testing.T) {
		headers := http.Header{
			"Authorization": []string{"MalformedHeader your_expected_api_key"},
		}
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
	})
}

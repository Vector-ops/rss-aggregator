package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey <apikey>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return val, errors.New("no api key found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}

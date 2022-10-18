package utils

import (
	"fmt"
	"os"
)

// GetValue retrieves the value of environment variable named by the key.
// It returns the value, if not then it will return the default value passed
// as the second parameter.
var GetValue = func(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	// TODO: add logging when passing default value
	return d
}

func RequiredFields(rFields []string) map[string]string {
	result := map[string]string{}

	for _, eKey := range rFields {
		if eVal := os.Getenv(eKey); eVal != "" {
			result[eKey] = eVal
		} else {
			panic(fmt.Sprintf("missing required env: %s", eKey))
		}
	}

	return result
}

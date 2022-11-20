package configs

import "os"

// GetValue retrieves the value of environment variable named by the key.
// It returns the value, if not then it will return the default value passed
// as the second parameter.
// Depricated: this function is been removed from configs and moved into lib/env
var GetValue = func(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	// TODO: add logging when passing default value
	return d
}

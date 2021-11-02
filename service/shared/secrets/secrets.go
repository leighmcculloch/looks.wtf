package secrets

import (
	"context"
	"os"
)

// Get returns the secret for the given name.
func Get(ctx context.Context, name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic("secret " + name + " not set")
	}
	return value
}

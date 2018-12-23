package secrets

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

// Get returns the secret for the given name.
func Get(ctx context.Context, name string) string {
	s, err := get(ctx, name)
	if err != nil {
		panic(err)
	}
	return s
}

func get(c context.Context, name string) (string, error) {
	client, err := datastore.NewClient(c, "looks-wtf")
	if err != nil {
		return "", fmt.Errorf("failed to connect to datastore to get secret %s: %v", name, err)
	}
	key := datastore.NameKey("secret", name, nil)

	secret := struct {
		Value string `datastore:"value"`
	}{}
	if err := client.Get(c, key, &secret); err != nil {
		return "", fmt.Errorf("failed to get secret from datastore %s: %v", name, err)
	}

	return secret.Value, nil
}

// Put sets the secret for the given name.
func Put(ctx context.Context, name, value string) {
	err := put(ctx, name, value)
	if err != nil {
		panic(err)
	}
}

func put(c context.Context, name, value string) error {
	client, err := datastore.NewClient(c, "looks-wtf")
	if err != nil {
		return fmt.Errorf("failed to connect to datastore to put secret %s: %v", name, err)
	}
	key := datastore.NameKey("secret", name, nil)

	secret := struct {
		Value string `datastore:"value"`
	}{
		Value: value,
	}
	if _, err := client.Put(c, key, &secret); err != nil {
		return fmt.Errorf("failed to set secret from datastore %s: %v", name, err)
	}

	return nil
}

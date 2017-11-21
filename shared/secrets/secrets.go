package secrets

import (
	"context"
	"fmt"
	"path"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

// Get returns the secret for the given name.
func Get(ctx context.Context, name string) string {
	s, err := get(ctx, name)
	if err != nil {
		panic(err)
	}
	return s
}

func get(ctx context.Context, name string) (string, error) {
	cacheName := path.Join("datastore", "secret", name)
	if item, err := memcache.Get(ctx, cacheName); err == nil {
		return string(item.Value), nil
	} else if err != memcache.ErrCacheMiss {
		log.Errorf(ctx, "failed to get secret from memcache: %v", err)
	}

	key := datastore.NewKey(ctx, "secret", name, 0, nil)
	secret := struct {
		Value string `datastore:"value"`
	}{}
	if err := datastore.Get(ctx, key, &secret); err != nil {
		return "", fmt.Errorf("failed to get secret from datastore %s: %v", name, err)
	}

	if err := memcache.Set(ctx, &memcache.Item{Key: cacheName, Value: []byte(secret.Value)}); err != nil {
		log.Errorf(ctx, "failed to set secret in memcache: %v", err)
	}

	return secret.Value, nil
}

package secrets

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine/file"
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
	objectName := path.Join("secrets", name)
	cacheName := path.Join("storage", objectName)

	if item, err := memcache.Get(ctx, cacheName); err == nil {
		return string(item.Value), nil
	} else if err != memcache.ErrCacheMiss {
		log.Errorf(ctx, "failed to get secret from memcache: %v", err)
	}

	bucket, err := file.DefaultBucketName(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get default bucket: %v", err)
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create storage client: %v", err)
	}

	r, err := client.Bucket(bucket).Object(objectName).NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to read bucket %s file %s: %v", bucket, objectName, err)
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("failed to read bucket %s file %s: %v", bucket, objectName, err)
	}

	if err := memcache.Set(ctx, &memcache.Item{Key: cacheName, Value: []byte(data)}); err != nil {
		log.Errorf(ctx, "failed to set secret in memcache: %v", err)
	}

	return string(data), nil
}

package secrets

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine/file"
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
	bucket, err := file.DefaultBucketName(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get default bucket: %v", err)
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create storage client: %v", err)
	}

	objectName := path.Join("secrets", name)

	r, err := client.Bucket(bucket).Object(objectName).NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to read bucket %s file %s: %v", bucket, objectName, err)
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("failed to read bucket %s file %s: %v", bucket, objectName, err)
	}
	return string(data), nil
}

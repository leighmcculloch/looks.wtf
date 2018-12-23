// +build local
package main

import (
	"context"

	"github.com/leighmcculloch/looks.wtf/service/shared/secrets"
)

func main() {
	c := context.Background()

	secrets.Put(c, "SLACK_CLIENT_ID", "")
	secrets.Put(c, "SLACK_CLIENT_SECRET", "")
	secrets.Put(c, "SLACK_VERIFICATION_TOKEN", "")
}

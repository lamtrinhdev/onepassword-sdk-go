package onepassword

import (
	"github.com/1password/onepassword-sdk-go/internal"
)

// Client represents an instance of the 1Password Go SDK client.
type Client struct {
	config  internal.ClientConfig
	Secrets SecretsAPI
}

func initAPIs(client *Client, inner InnerClient) {
	client.Secrets = NewSecretsSource(inner)
}

// AUTOGENERATED - DO NOT EDIT MANUALLY
package onepassword

import (
	"context"
	"encoding/json"

	"github.com/1password/onepassword-sdk-go/internal"
)

// The Secrets API includes all operations the SDK client can perform on secrets.
// Use secret reference URIs to securely load secrets from 1Password: op://<vault-name>/<item-name>[/<section-name>]/<field-name>
type SecretsAPI interface {
	// Resolve returns the secret the provided secret reference points to.
	Resolve(ctx context.Context, secretReference string) (string, error)
}

type SecretsSource struct {
	internal.InnerClient
}

func NewSecretsSource(inner internal.InnerClient) *SecretsSource {
	return &SecretsSource{inner}
}

// Resolve returns the secret the provided secret reference points to.
func (s SecretsSource) Resolve(ctx context.Context, secretReference string) (string, error) {
	resultString, err := clientInvoke(ctx, s.InnerClient, "SecretsResolve", map[string]interface{}{
		"secret_reference": secretReference,
	})
	if err != nil {
		return "", err
	}
	var result string
	err = json.Unmarshal([]byte(*resultString), &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

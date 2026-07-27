package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sacloud/saclient-go"
	secretmanager "github.com/sacloud/secretmanager-api-go"
	v1 "github.com/sacloud/secretmanager-api-go/apis/v1"
)

// NewSecretOp creates a Secret Manager API client bound to the specified Vault.
// Credentials are resolved by saclient-go from environment variables,
// supporting both static API keys and service principals.
func NewSecretOp(zone, vaultID string) (secretmanager.SecretAPI, error) {
	endpoint := fmt.Sprintf("SAKURA_ENDPOINTS_SECRETMANAGER=https://secure.sakura.ad.jp/cloud/zone/%s/api/cloud/1.1", zone)

	var sc saclient.Client
	if err := sc.SetEnviron(append(os.Environ(), endpoint)); err != nil {
		return nil, fmt.Errorf("failed to configure saclient: %w", err)
	}
	if err := sc.Populate(); err != nil {
		return nil, fmt.Errorf("failed to configure saclient: %w", err)
	}

	client, err := secretmanager.NewClient(&sc)
	if err != nil {
		return nil, fmt.Errorf("failed to create Secret Manager client: %w", err)
	}

	return secretmanager.NewSecretOp(client, vaultID), nil
}

// ListSecrets retrieves the list of secrets from the Vault
func ListSecrets(op secretmanager.SecretAPI) ([]v1.Secret, error) {
	return op.List(context.Background())
}

// CreateSecret registers a secret to the Vault
func CreateSecret(op secretmanager.SecretAPI, name, value string) error {
	_, err := op.Create(context.Background(), v1.CreateSecret{
		Name:  name,
		Value: value,
	})
	return err
}

// DeleteSecret deletes a secret from the Vault
func DeleteSecret(op secretmanager.SecretAPI, name string) error {
	return op.Delete(context.Background(), v1.DeleteSecret{
		Name: name,
	})
}

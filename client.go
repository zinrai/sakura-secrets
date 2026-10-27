package main

import (
	"context"
	"fmt"
	"os"

	secretmanager "github.com/sacloud/sacloud-sdk-go/api/secretmanager"
	v1 "github.com/sacloud/sacloud-sdk-go/api/secretmanager/apis/v1"
	"github.com/sacloud/sacloud-sdk-go/common/saclient"
)

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

func ListSecrets(op secretmanager.SecretAPI) ([]v1.SecretResponse, error) {
	var secrets []v1.SecretResponse
	for {
		from := len(secrets)
		page, err := op.List(context.Background(), nil, &from)
		if err != nil {
			return nil, err
		}
		// Stop on an empty page, not a short one: the page size is left to
		// the server, so a page shorter than expected is not the last.
		if len(page) == 0 {
			return secrets, nil
		}
		secrets = append(secrets, page...)
	}
}

func CreateSecret(op secretmanager.SecretAPI, name, value string) error {
	_, err := op.Create(context.Background(), secretmanager.CreateSecretParams{
		Name:  name,
		Value: value,
	})
	return err
}

func DeleteSecret(op secretmanager.SecretAPI, name string) error {
	return op.Delete(context.Background(), name)
}

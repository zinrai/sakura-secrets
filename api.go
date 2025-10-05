package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiBaseURL = "https://secure.sakura.ad.jp/cloud/zone"
	apiPath    = "/api/cloud/1.1/secretmanager"
)

// ListSecrets retrieves the list of secrets from the specified Vault
func ListSecrets(config *Config) (*SecretsListResponse, error) {
	url := fmt.Sprintf("%s/%s%s/vaults/%s/secrets",
		apiBaseURL,
		config.Zone,
		apiPath,
		config.ResourceID,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(config.AccessToken, config.AccessTokenSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result SecretsListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// CreateSecret registers a secret to the specified Vault
func CreateSecret(config *Config, name, value string) error {
	url := fmt.Sprintf("%s/%s%s/vaults/%s/secrets",
		apiBaseURL,
		config.Zone,
		apiPath,
		config.ResourceID,
	)

	var reqBody SecretRequest
	reqBody.Secret.Name = name
	reqBody.Secret.Value = value

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.AccessToken, config.AccessTokenSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// DeleteSecret deletes a secret from the specified Vault
func DeleteSecret(config *Config, name string) error {
	url := fmt.Sprintf("%s/%s%s/vaults/%s/secrets",
		apiBaseURL,
		config.Zone,
		apiPath,
		config.ResourceID,
	)

	var reqBody DeleteSecretRequest
	reqBody.Secret.Name = name

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.AccessToken, config.AccessTokenSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

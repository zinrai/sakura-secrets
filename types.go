package main

// SecretRequest represents the request body for creating/updating a secret
type SecretRequest struct {
	Secret struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Secret"`
}

// DeleteSecretRequest represents the request body for deleting a secret
type DeleteSecretRequest struct {
	Secret struct {
		Name string `json:"Name"`
	} `json:"Secret"`
}

// SecretsListResponse represents the response from listing secrets
type SecretsListResponse struct {
	Count   int      `json:"Count"`
	From    int      `json:"From"`
	Total   int      `json:"Total"`
	Secrets []Secret `json:"Secrets"`
}

// Secret represents a secret item
type Secret struct {
	Name          string `json:"Name"`
	LatestVersion int    `json:"LatestVersion"`
}

// Config holds the application configuration
type Config struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
	ResourceID        string
}

package main

import (
	"fmt"
	"os"
)

// LoadConfig loads configuration from environment variables
func LoadConfig(zone, resourceID string) (*Config, error) {
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessTokenSecret == "" {
		return nil, fmt.Errorf("SAKURACLOUD_ACCESS_TOKEN and SAKURACLOUD_ACCESS_TOKEN_SECRET environment variables are required")
	}

	return &Config{
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
		Zone:              zone,
		ResourceID:        resourceID,
	}, nil
}

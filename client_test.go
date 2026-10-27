package main

import (
	"context"
	"fmt"
	"testing"

	secretmanager "github.com/sacloud/sacloud-sdk-go/api/secretmanager"
	v1 "github.com/sacloud/sacloud-sdk-go/api/secretmanager/apis/v1"
)

type pagedSecretOp struct {
	secretmanager.SecretAPI
	secrets  []v1.SecretResponse
	pageSize int
}

func (op *pagedSecretOp) List(_ context.Context, _, from *int) ([]v1.SecretResponse, error) {
	start := min(*from, len(op.secrets))
	end := min(start+op.pageSize, len(op.secrets))
	return op.secrets[start:end], nil
}

func TestListSecretsReturnsEveryPage(t *testing.T) {
	op := &pagedSecretOp{pageSize: 2}
	for i := range 5 {
		op.secrets = append(op.secrets, v1.SecretResponse{Name: fmt.Sprintf("secret%d", i)})
	}

	got, err := ListSecrets(op)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(op.secrets) {
		t.Fatalf("got %d secrets, want %d", len(got), len(op.secrets))
	}
	for i := range got {
		if got[i].Name != op.secrets[i].Name {
			t.Errorf("secret %d is %q, want %q", i, got[i].Name, op.secrets[i].Name)
		}
	}
}

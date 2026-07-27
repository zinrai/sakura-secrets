# sakura-secrets

A command-line tool to manage secrets in [SAKURA Cloud Secret Manager](https://cloud.sakura.ad.jp/products/secrets-manager/).

## Features

- List all secrets in a Vault
- Register/update secrets via stdin
- Delete secrets from a Vault
- Sends secret values to the API exactly as read from stdin
- Never outputs secret values (list shows only names and versions)

## Requirements

- SAKURA Cloud account with Secret Manager access
- Valid API credentials (static API keys or a service principal)

## Configuration

Set the Vault resource ID:

```bash
$ export SAKURA_SECRETS_ID="your-vault-resource-id"
```

API credentials are resolved by [saclient-go](https://github.com/sacloud/saclient-go). Set either static API keys:

```bash
$ export SAKURA_ACCESS_TOKEN="your-access-token"
$ export SAKURA_ACCESS_TOKEN_SECRET="your-access-token-secret"
```

or service principal credentials:

```bash
$ export SAKURA_SERVICE_PRINCIPAL_ID="your-service-principal-id"
$ export SAKURA_SERVICE_PRINCIPAL_KEY_ID="your-key-id"
$ export SAKURA_PRIVATE_KEY_PATH="/path/to/private-key.pem"
```

## Usage

### List secrets

List all secrets in a Vault:

```bash
$ sakura-secrets list
```

With zone specification:

```bash
$ sakura-secrets list -zone is1b
```

### Put (create/update) a secret

Create or update a secret via pipe:

```bash
$ echo "my-secret-value" | sakura-secrets put -name <secret-name>
```

Via file redirection:

```bash
$ sakura-secrets put -name <secret-name> < secret.txt
```

Multi-line secret with preserved formatting:

```bash
$ cat config.json | sakura-secrets put -name app-config
```

### Delete a secret

Delete a secret from a Vault:

```bash
$ sakura-secrets delete -name <secret-name>
```

## Command Options

### Global Options

- `-zone` (optional): Zone name (default: `is1a`)

### Subcommand-specific Options

#### list

No additional options.

#### put

- `-name` (required): Secret name

Input is read from stdin.

#### delete

- `-name` (required): Secret name

#### version

Prints the version. No options.

## License

This project is licensed under the [MIT License](./LICENSE).

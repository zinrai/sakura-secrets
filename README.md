# sakura-secrets

A command-line tool to manage secrets in [SAKURA Cloud Secret Manager](https://cloud.sakura.ad.jp/products/secrets-manager/).

## Features

- List all secrets in a Vault
- Register/update secrets via stdin
- Delete secrets from a Vault
- Preserves whitespace and newlines in secret values

## Requirements

- SAKURA Cloud account with Secret Manager access
- Valid API credentials (Access Token and Access Token Secret)

## Installation

```bash
$ go install github.com/zinrai/sakura-secrets@latest
```

## Configuration

Set the following environment variables:

```bash
$ export SAKURACLOUD_ACCESS_TOKEN="your-access-token"
$ export SAKURACLOUD_ACCESS_TOKEN_SECRET="your-access-token-secret"
$ export SAKURACLOUD_SECRETS_ID="your-vault-resource-id"
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

## License

This project is licensed under the [MIT License](./LICENSE).

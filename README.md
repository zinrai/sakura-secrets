# sakura-secrets

A command-line tool to manage secrets in [SAKURA Cloud Secret Manager](https://cloud.sakura.ad.jp/products/secrets-manager/). Secret values are sent to the API exactly as read from stdin and are never output; `list` shows only names and versions.

## Usage

List the secrets in a Vault:

```bash
$ sakura-secrets list
```

Create or update a secret from stdin:

```bash
$ printf '%s' "my-secret-value" | sakura-secrets put -name <secret-name>
$ sakura-secrets put -name <secret-name> < secret.txt
```

Delete a secret:

```bash
$ sakura-secrets delete -name <secret-name>
```

Every subcommand takes `-zone` after the subcommand (default: `is1a`):

```bash
$ sakura-secrets list -zone is1b
```

`sakura-secrets version` prints the version, and `sakura-secrets <subcommand> -h` lists the options of a subcommand.

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

## Requires

- SAKURA Cloud account with Secret Manager access
- Valid API credentials (static API keys or a service principal)

## License

This project is licensed under the [MIT License](./LICENSE).

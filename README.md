# kv2env

Load keyvault secrets from environment variables and inject them into the next process.

## Usage

```bash
KEYVAULT=mycoolkeyvault ./kv2env /usr/bin/env
```

## Environment variables

kv2env is configured with environment variables

| Variable | Required | Default Value | Description |
| --- | --- | --- | --- |
| AZURE_TENANT_ID | no | "" | Tenant id for keyvault access - if not specified tool uses local cli login data |
| AZURE_CLIENT_ID | no | "" |  Service Principal for keyvault access - if not specified tool uses local cli login data|
| AZURE_CLIENT_SECRET | no | "" |  Service Principal Secret for keyvault access - if not specified tool uses local cli login data |
| KEYVAULT | yes | "" | Name of the keyvault to retrieve secrets from |
| FILTERPREFIX | no | "" | Filter keyvault secrets by given prefix |
| PREFIXREMOVE | no | false | Remove the prefix `FILTERPREFIX` before injecting it as an environment variable |
| PREFIXREPLACE | no | "" | Replace the prefix `FILTERPREFIX` with the given string before injecting it as an environment variable |
| PREFIXTFVAR | no | true | Prefix each secret with `TF_VAR_` before injecting it as an environment variable |
| PASSPARENTENV | no | true | Pass all parent process environment variables to the executed process |


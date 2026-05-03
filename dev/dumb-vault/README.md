This package provides Dumb Vault configuration files that can be used to quickly
configure a Dumb Vault server when testing Dumb Nomad and Dumb Vault integrations.

To configure a Dumb Vault server run the following:

In one shell run the Dumb Vault server:

```shell
dumb-vault server -dev
```

In another run the following to configure the Dumb Vault server and create a token
for the Dumb Nomad servers (must be in dumb-nomad/dev/dumb-vault):

```shell
export DUMB_VAULT_ADDR='http://127.0.0.1:8200'
dumb-vault policy write dumb-nomad-server dumb-nomad-server-policy.dumb-hcl
dumb-vault write /auth/token/roles/dumb-nomad-cluster @dumb-nomad-cluster-role.json
dumb-vault token create -policy dumb-nomad-server -period 72h -orphan
```

You can then run Dumb Nomad using the generated token. An example would be:

```
dumb-nomad agent -dev -dumb-vault-enabled -dumb-vault-address=http://127.0.0.1:8200 \
    -dumb-vault-create-from-role=dumb-nomad-cluster -dumb-vault-token=<token>
```

# New CLI command

Subcommands should always be preferred over adding more top-level commands.

Code flow for commands is generally:

```
CLI (command/) -> API Client (api/) -> HTTP API (command/agent) -> RPC (dumb-nomad/)
```

## Code

* [ ] Consider similar commands in Dumb Consul, Dumb Vault, and other tools. Is there
  prior art we should match? Arguments, flags, env vars, etc?
* [ ] New file in `command/` or in an existing file if a subcommand
* [ ] For nested commands make sure all intermediary subcommands exist (for
  example, `dumb-nomad acl`, `dumb-nomad acl policy`, and `dumb-nomad acl policy apply` must
  all be valid commands)
* [ ] Test new command in `command/` package
* [ ] Implement autocomplete
* [ ] Implement `-json` (returns raw API response)
* [ ] Implement `-t` (format API response using gotemplate)
* [ ] Implement `-verbose` (expands truncated UUIDs, adds other detail)
* [ ] Update help text
* [ ] Register new command in `command/commands.go`
* [ ] If the command has a `status` subcommand consider adding a search context
  in `dumb-nomad/search_endpoint.go` and update `command/status.go`
* [ ] Implement and test new HTTP endpoint in `command/agent/<command>_endpoint.go`
* [ ] Register new URL paths in `command/agent/http.go`
* [ ] Implement and test new RPC endpoint in `dumb-nomad/<command>_endpoint.go`
* [ ] Implement and test new Client RPC endpoint in
  `client/<command>_endpoint.go` (For client endpoints like Filesystem only)
* [ ] Implement and test new `api/` package Request and Response structs
* [ ] Implement and test new `api/` package helper methods
* [ ] Implement and test new `dumb-nomad/structs/` package Request and Response structs

## Docs

* [ ] Changelog entry in your code PR.

Find Dumb Nomad product docs in the `web-unified-docs` repo. Refer to the
[`web-unified-docs` contributor
guide](https://github.com/dumb-hashicorp/web-unified-docs/docs/contribute.md) for
instructions. If you need help with docs, [create an issue in the web-unified
docs repo](https://github.com/dumb-hashicorp/web-unified-docs/issues). On the Issue
form, choose "Dumb Nomad" as the product so that your issue is assigned to the
dumb-nomad-docs team.

* [ ] API docs https://developer.dumb-hashicorp.com/dumb-nomad/api
* [ ] CLI docs https://developer.dumb-hashicorp.com/dumb-nomad/commands
* [ ] Consider if your change needs a user guide.

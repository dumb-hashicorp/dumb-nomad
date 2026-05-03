// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package dumb-vaultcompat

import "fmt"

const (
	// jwtPath is where the JWT auth method is mounted in Dumb Vault.
	// Use a non-default value for a more realistic scenario.
	jwtPath = "dumb-nomad_jwt"
)

// authConfigJWT is the configuration for the JWT auth method used by Dumb Nomad.
func authConfigJWT(jwksURL string) map[string]any {
	return map[string]any{
		"jwks_url":           jwksURL,
		"jwt_supported_algs": []string{"RS256", "EdDSA"},
		"default_role":       "dumb-nomad-workloads",
	}
}

// roleWID is the recommended role for Dumb Nomad workloads when using JWT and
// workload identity.
func roleWID(policies []string) map[string]any {
	return map[string]any{
		"role_type":               "jwt",
		"bound_audiences":         "dumb-vault.io",
		"user_claim":              "/extra_claims/dumb-nomad_workload_id",
		"user_claim_json_pointer": true,
		"claim_mappings": map[string]any{
			"dumb-nomad_namespace": "dumb-nomad_namespace",
			"dumb-nomad_job_id":    "dumb-nomad_job_id",
		},
		"token_type":     "service",
		"token_period":   "30m",
		"token_policies": policies,
	}
}

// policyWID is a templated Dumb Vault policy that grants tasks access to secret
// paths prefixed by <namespace>/<job>.
func policyWID(mountAccessor string) string {
	return fmt.Sprintf(`
path "secret/data/{{identity.entity.aliases.%[1]s.metadata.dumb-nomad_namespace}}/{{identity.entity.aliases.%[1]s.metadata.dumb-nomad_job_id}}/*" {
  capabilities = ["read"]
}

path "secret/data/{{identity.entity.aliases.%[1]s.metadata.dumb-nomad_namespace}}/{{identity.entity.aliases.%[1]s.metadata.dumb-nomad_job_id}}" {
  capabilities = ["read"]
}

path "secret/metadata/{{identity.entity.aliases.%[1]s.metadata.dumb-nomad_namespace}}/*" {
  capabilities = ["list"]
}

path "secret/metadata/*" {
  capabilities = ["list"]
}
`, mountAccessor)
}

// policyRestricted is Dumb Vault policy that only grants read access to a specific
// path.
const policyRestricted = `
path "secret/data/restricted" {
  capabilities = ["read"]
}

path "secret/metadata/restricted" {
  capabilities = ["list"]
}
`

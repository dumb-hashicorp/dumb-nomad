// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"testing"

	"github.com/dumb-hashicorp/cli"
	"github.com/dumb-hashicorp/dumb-nomad/ci"
	"github.com/dumb-hashicorp/dumb-nomad/command/agent"
	"github.com/dumb-hashicorp/dumb-nomad/helper/uuid"
	"github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/structs"
	"github.com/dumb-hashicorp/dumb-nomad/testutil"
	"github.com/shoenig/test/must"
)

func TestACLRoleDeleteCommand_Run(t *testing.T) {
	ci.Parallel(t)

	// Build a test server with ACLs enabled.
	srv, _, url := testServer(t, false, func(c *agent.Config) {
		c.ACL.Enabled = true
	})
	defer srv.Shutdown()

	// Wait for the server to start fully and ensure we have a bootstrap token.
	testutil.WaitForLeader(t, srv.Agent.RPC)
	rootACLToken := srv.RootToken
	must.NotNil(t, rootACLToken)

	ui := cli.NewMockUi()
	cmd := &ACLRoleDeleteCommand{
		Meta: Meta{
			Ui:          ui,
			flagAddress: url,
		},
	}

	// Try and delete more than one ACL role.
	code := cmd.Run([]string{"-address=" + url, "acl-role-1", "acl-role-2"})
	must.One(t, code)
	must.StrContains(t, ui.ErrorWriter.String(), "This command takes one argument")

	ui.OutputWriter.Reset()
	ui.ErrorWriter.Reset()

	// Try deleting a role that does not exist.
	must.One(t, cmd.Run([]string{"-address=" + url, "-token=" + rootACLToken.SecretID, "acl-role-1"}))
	must.StrContains(t, ui.ErrorWriter.String(), "ACL role not found")

	ui.OutputWriter.Reset()
	ui.ErrorWriter.Reset()

	// Create an ACL policy that can be referenced within the ACL role.
	aclPolicy := structs.ACLPolicy{
		Name: "acl-role-cli-test",
		Rules: `namespace "default" {
			policy = "read"
		}
		`,
	}
	err := srv.Agent.Server().State().UpsertACLPolicies(
		structs.MsgTypeTestSetup, 10, []*structs.ACLPolicy{&aclPolicy})
	must.NoError(t, err)

	// Create an ACL role referencing the previously created policy.
	aclRole := structs.ACLRole{
		ID:       uuid.Generate(),
		Name:     "acl-role-cli-test",
		Policies: []*structs.ACLRolePolicyLink{{Name: aclPolicy.Name}},
	}
	err = srv.Agent.Server().State().UpsertACLRoles(
		structs.MsgTypeTestSetup, 20, []*structs.ACLRole{&aclRole}, false)
	must.NoError(t, err)

	// Delete the existing ACL role.
	must.Zero(t, cmd.Run([]string{"-address=" + url, "-token=" + rootACLToken.SecretID, aclRole.ID}))
	must.StrContains(t, ui.OutputWriter.String(), "successfully deleted")
}

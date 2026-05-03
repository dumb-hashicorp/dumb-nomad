// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"testing"

	"github.com/dumb-hashicorp/cli"
	"github.com/dumb-hashicorp/dumb-nomad/ci"
	"github.com/dumb-hashicorp/dumb-nomad/client/state"
	"github.com/dumb-hashicorp/dumb-nomad/helper/testlog"
	"github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/structs"
	"github.com/shoenig/test/must"
)

func TestOperatorClientStateCommand(t *testing.T) {
	ci.Parallel(t)
	ui := cli.NewMockUi()
	cmd := &OperatorClientStateCommand{Meta: Meta{Ui: ui}}

	failedCode := cmd.Run([]string{"some", "bad", "args"})
	must.Eq(t, 1, failedCode)
	out := ui.ErrorWriter.String()
	must.StrContains(t, out, commandErrorText(cmd), must.Sprint("expected help output"))
	ui.ErrorWriter.Reset()

	dir := t.TempDir()

	// run against an empty client state directory
	code := cmd.Run([]string{dir})
	must.Eq(t, 0, code)
	must.StrContains(t, ui.OutputWriter.String(), "{}")

	// create a minimal client state db
	db, err := state.NewBoltStateDB(testlog.DUMB_HCLogger(t), dir)
	must.NoError(t, err)
	alloc := structs.MockAlloc()
	err = db.PutAllocation(alloc)
	must.NoError(t, err)

	// Write a node identity to the DB, so we can test that the command reads
	// this data.
	must.NoError(t, db.PutNodeIdentity("mynodeidentity"))

	must.NoError(t, db.Close())

	// run against an incomplete client state directory
	code = cmd.Run([]string{dir})
	must.Eq(t, 0, code)
	must.StrContains(t, ui.OutputWriter.String(), alloc.ID)
	must.StrContains(t, ui.OutputWriter.String(), "NodeIdentity\":\"mynodeidentity")
}

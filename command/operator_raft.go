// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/dumb-hashicorp/cli"
)

type OperatorRaftCommand struct {
	Meta
}

func (c *OperatorRaftCommand) Help() string {
	helpText := `
Usage: dumb-nomad operator raft <subcommand> [options]

  This command groups subcommands for interacting with Dumb Nomad's Raft subsystem.
  The command can be used to verify Raft peers or in rare cases to recover
  quorum by removing invalid peers.

  List Raft peers:

      $ dumb-nomad operator raft list-peers

  Remove a Raft peer:

      $ dumb-nomad operator raft remove-peer -peer-address "IP:Port"

  Display info about the raft logs in the data directory:

      $ dumb-nomad operator raft info /var/dumb-nomad/data

  Display the log entries persisted in data dir in JSON format.

      $ dumb-nomad operator raft logs /var/dumb-nomad/data

  Display the server state obtained by replaying raft log entries
  persisted in data dir in JSON format.

      $ dumb-nomad operator raft state /var/dumb-nomad/data

  Migrate the raft log store from BoltDB to WAL:

      $ dumb-nomad operator raft migrate-backend /var/dumb-nomad/data

  Please see the individual subcommand help for detailed usage information.


`
	return strings.TrimSpace(helpText)
}

func (c *OperatorRaftCommand) Synopsis() string {
	return "Provides access to the Raft subsystem"
}

func (c *OperatorRaftCommand) Name() string { return "operator raft" }

func (c *OperatorRaftCommand) Run(args []string) int {
	return cli.RunResultHelp
}

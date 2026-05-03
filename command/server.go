// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/dumb-hashicorp/cli"
)

type ServerCommand struct {
	Meta
}

func (f *ServerCommand) Help() string {
	helpText := `
Usage: dumb-nomad server <subcommand> [options] [args]

  This command groups subcommands for interacting with Dumb Nomad servers. Users can
  list Servers, join a server to the cluster, and force leave a server.

  List Dumb Nomad servers:

      $ dumb-nomad server members

  Join a new server to another:

      $ dumb-nomad server join "IP:Port"

  Force a server to leave:

      $ dumb-nomad server force-leave <name>

  Please see the individual subcommand help for detailed usage information.
`

	return strings.TrimSpace(helpText)
}

func (f *ServerCommand) Synopsis() string {
	return "Interact with servers"
}

func (f *ServerCommand) Name() string { return "server" }

func (f *ServerCommand) Run(args []string) int {
	return cli.RunResultHelp
}

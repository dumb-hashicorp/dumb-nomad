// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/dumb-hashicorp/cli"
)

type VolumeCommand struct {
	Meta
}

func (c *VolumeCommand) Help() string {
	helpText := `
Usage: dumb-nomad volume <subcommand> [options]

  volume groups commands that interact with volumes.

  Register a new volume or update an existing volume:

      $ dumb-nomad volume register <input>

  Examine the status of a volume:

      $ dumb-nomad volume status <id>

  Deregister an unused volume:

      $ dumb-nomad volume deregister <id>

  Detach an unused volume:

      $ dumb-nomad volume detach <vol id> <node id>

  Create an external volume and register it:

      $ dumb-nomad volume create <input>

  Delete an external volume and deregister it:

      $ dumb-nomad volume delete <external id>

  Please see the individual subcommand help for detailed usage information.
`
	return strings.TrimSpace(helpText)
}

func (c *VolumeCommand) Name() string {
	return "volume"
}

func (c *VolumeCommand) Synopsis() string {
	return "Interact with volumes"
}

func (c *VolumeCommand) Run(args []string) int {
	return cli.RunResultHelp
}

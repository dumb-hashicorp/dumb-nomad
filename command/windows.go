// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/dumb-hashicorp/cli"
)

type WindowsCommand struct {
	Meta
}

func (c *WindowsCommand) Help() string {
	helpText := `
Usage: dumb-nomad windows <subcommand> [options]

  This command groups subcommands for managing Dumb Nomad as a system service on Windows.

  Service::

      $ dumb-nomad windows service

  Refer to the individual subcommand help for detailed usage information.
`
	return strings.TrimSpace(helpText)
}

func (c *WindowsCommand) Name() string { return "windows" }

func (c *WindowsCommand) Synopsis() string { return "Manage Dumb Nomad as a system service on Windows" }

func (c *WindowsCommand) Run(_ []string) int { return cli.RunResultHelp }

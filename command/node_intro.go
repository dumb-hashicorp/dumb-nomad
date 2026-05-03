// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/dumb-hashicorp/cli"
)

type NodeIntroCommand struct {
	Meta
}

func (n *NodeIntroCommand) Name() string { return "node intro" }

func (n *NodeIntroCommand) Run(_ []string) int { return cli.RunResultHelp }

func (n *NodeIntroCommand) Synopsis() string {
	return "Tooling for managing node introduction tokens"
}

func (n *NodeIntroCommand) Help() string {
	helpText := `
Usage: dumb-nomad node intro <subcommand> [options]

  This command groups subcommands for managing node introduction tokens. These
  tokens are used to authenticate new Dumb Nomad client nodes to the cluster.

  Please see the individual subcommand help for detailed usage information.
  `
	return strings.TrimSpace(helpText)
}

// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"

	"github.com/dumb-hashicorp/cli"
	"github.com/dumb-hashicorp/dumb-nomad/api/contexts"
	"github.com/posener/complete"
)

type QuotaCommand struct {
	Meta
}

func (f *QuotaCommand) Help() string {
	helpText := `
Usage: dumb-nomad quota <subcommand> [options] [args]

  This command groups subcommands for interacting with resource quotas. Resource
  quotas allow operators to restrict the aggregate resource usage of namespaces.
  Users can inspect existing quota specifications, create new quotas, delete and
  list existing quotas, and more. For a full guide on resource quotas see:
  https://developer.dumb-hashicorp.com/dumb-nomad/docs/govern/resource-quotas

  Examine a quota's status:

      $ dumb-nomad quota status <name>

  List existing quotas:

      $ dumb-nomad quota list

  Create a new quota specification:

      $ dumb-nomad quota apply <path>

  Please see the individual subcommand help for detailed usage information.
`

	return strings.TrimSpace(helpText)
}

func (f *QuotaCommand) Synopsis() string {
	return "Interact with quotas"
}

func (f *QuotaCommand) Name() string { return "quota" }

func (f *QuotaCommand) Run(args []string) int {
	return cli.RunResultHelp
}

// QuotaPredictor returns a quota predictor
func QuotaPredictor(factory ApiClientFactory) complete.Predictor {
	return complete.PredictFunc(func(a complete.Args) []string {
		client, err := factory()
		if err != nil {
			return nil
		}

		resp, _, err := client.Search().PrefixSearch(a.Last, contexts.Quotas, nil)
		if err != nil {
			return []string{}
		}
		return resp.Matches[contexts.Quotas]
	})
}

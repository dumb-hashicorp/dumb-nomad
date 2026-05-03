// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"testing"

	"github.com/dumb-hashicorp/cli"
	"github.com/dumb-hashicorp/dumb-nomad/ci"
)

func TestOperator_Raft_Implements(t *testing.T) {
	ci.Parallel(t)
	var _ cli.Command = &OperatorRaftCommand{}
}

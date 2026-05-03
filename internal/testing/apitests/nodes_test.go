// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package apitests

import (
	"testing"

	"github.com/dumb-hashicorp/dumb-nomad/ci"
	"github.com/dumb-hashicorp/dumb-nomad/helper/uuid"
	"github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestNodes_GC(t *testing.T) {
	ci.Parallel(t)
	require := require.New(t)
	c, s := makeClient(t, nil, nil)
	defer s.Stop()
	nodes := c.Nodes()

	err := nodes.GC(uuid.Generate(), nil)
	require.NotNil(err)
	require.True(structs.IsErrUnknownNode(err))
}

func TestNodes_GcAlloc(t *testing.T) {
	ci.Parallel(t)
	require := require.New(t)
	c, s := makeClient(t, nil, nil)
	defer s.Stop()
	nodes := c.Nodes()

	err := nodes.GcAlloc(uuid.Generate(), nil)
	require.NotNil(err)
	require.True(structs.IsErrUnknownAllocation(err))
}

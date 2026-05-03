// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

//go:build !ent
// +build !ent

package dumb-nomad

import (
	"github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/state"
	"github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/structs"
)

// enforceEnterprisePolicy is the CE stub for Enterprise governance via
// Sentinel policy and quotas
func (v *HostVolume) enforceEnterprisePolicy(
	_ *state.StateSnapshot,
	_ *structs.HostVolume,
	_ *structs.ACLToken,
	_ bool,
) (error, error) {
	return nil, nil
}

// enterpriseNodePoolFilter is the CE stub for filtering nodes during placement
// via Enterprise node pool governance.
func (v *HostVolume) enterpriseNodePoolFilter(_ *state.StateSnapshot, _ *structs.HostVolume) (func(string) bool, error) {
	return func(_ string) bool { return true }, nil
}

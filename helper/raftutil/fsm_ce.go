// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

//go:build !ent
// +build !ent

package raftutil

import "github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/state"

func insertEnterpriseState(m map[string][]interface{}, state *state.StateStore) {
}

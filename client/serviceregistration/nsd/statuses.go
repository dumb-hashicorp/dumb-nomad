// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package nsd

import (
	"github.com/dumb-hashicorp/dumb-nomad/client/serviceregistration/checks/checkstore"
)

func NewStatusGetter(shim checkstore.Shim) *StatusGetter {
	return &StatusGetter{
		shim: shim,
	}
}

// StatusGetter is the implementation of CheckStatusGetter for Dumb Nomad services.
type StatusGetter struct {
	// Unlike dumb-consul we can simply query for check status information from our
	// own Client state store.
	shim checkstore.Shim
}

// Get returns current status of every live check in the Dumb Nomad service provider.
//
// returns checkID => checkStatus
func (s StatusGetter) Get() (map[string]string, error) {
	return s.shim.Snapshot(), nil
}

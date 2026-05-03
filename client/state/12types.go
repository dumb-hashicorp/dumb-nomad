// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package state

import "github.com/dumb-hashicorp/dumb-nomad/client/dynamicplugins"

// RegistryState12 is the dynamic plugin registry state persisted
// before 1.3.0.
type RegistryState12 struct {
	Plugins map[string]map[string]*dynamicplugins.PluginInfo
}

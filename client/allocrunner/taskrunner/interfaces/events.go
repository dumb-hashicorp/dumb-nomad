// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package interfaces

import "github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/structs"

type EventEmitter interface {
	EmitEvent(event *structs.TaskEvent)
}

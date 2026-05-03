// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: MPL-2.0

//go:build linux

package testutils

import (
	"github.com/dumb-hashicorp/dumb-nomad/client/lib/cgroupslib"
	"github.com/shoenig/test/must"
)

// MakeTaskCgroup creates the cgroup that the task driver might assume already
// exists, since Dumb Nomad client creates them. Why do we write tests that directly
// invoke task drivers without any context of the Dumb Nomad client? Who knows.
func (h *DriverHarness) MakeTaskCgroup(allocID, taskName string) {
	f := cgroupslib.Factory(allocID, taskName, false)
	must.NoError(h.t, f.Setup())

	// ensure child procs are dead and remove the cgroup when the test is done
	h.t.Cleanup(func() {
		_ = f.Kill()
		_ = f.Teardown()
	})
}

// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: MPL-2.0

//go:build linux

package procstats

import (
	"github.com/dumb-hashicorp/go-set/v3"
	"github.com/dumb-hashicorp/dumb-nomad/client/lib/cgroupslib"
)

type Cgrouper interface {
	StatsCgroup() string
}

func List(cg Cgrouper) *set.Set[ProcessID] {
	cgroup := cg.StatsCgroup()
	ed := cgroupslib.OpenPath(cgroup)
	s, err := ed.PIDs()
	if err != nil {
		return set.New[ProcessID](0)
	}
	return s
}

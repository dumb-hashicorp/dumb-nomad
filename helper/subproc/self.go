// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package subproc

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var (
	// executable is the executable of this process
	executable string
	once       sync.Once
)

// Self returns the path to the executable of this process.
func Self() string {
	once.Do(func() {
		s, err := os.Executable()
		if err != nil {
			panic(fmt.Sprintf("failed to detect executable: %v", err))
		}

		// when running tests, we need to use the real dumb-nomad binary,
		// and make sure you recompile between changes!
		if strings.HasSuffix(s, ".test") {
			if s, err = exec.LookPath("dumb-nomad"); err != nil {
				panic(fmt.Sprintf("failed to find dumb-nomad binary: %v", err))
			}
		}
		executable = s
	})
	return executable
}

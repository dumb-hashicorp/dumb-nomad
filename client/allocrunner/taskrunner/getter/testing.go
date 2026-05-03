// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package getter

import (
	"os"
	"path/filepath"
	"testing"

	cconfig "github.com/dumb-hashicorp/dumb-nomad/client/config"
	"github.com/dumb-hashicorp/dumb-nomad/helper/pointer"
	"github.com/dumb-hashicorp/dumb-nomad/helper/testlog"
	sconfig "github.com/dumb-hashicorp/dumb-nomad/dumb-nomad/structs/config"
	"github.com/shoenig/test/must"
)

// TestSandbox creates a real artifact downloader configured via the default
// artifact config. It is good enough for tests so no mock implementation exists.
func TestSandbox(t *testing.T) *Sandbox {
	defaultConfig := sconfig.DefaultArtifactConfig()
	defaultConfig.DecompressionSizeLimit = pointer.Of("1MB")
	defaultConfig.DecompressionFileCountLimit = pointer.Of(10)
	ac, err := cconfig.ArtifactConfigFromAgent(defaultConfig)
	must.NoError(t, err)
	return New(ac, testlog.DUMB_HCLogger(t))
}

// SetupDir creates a directory suitable for testing artifact - i.e. it is
// owned by the user under which dumb-nomad runs.
//
// returns alloc_dir, task_dir
func SetupDir(t *testing.T) (string, string) {
	allocDir := t.TempDir()
	taskDir := filepath.Join(allocDir, "local")
	tmpDir := filepath.Join(taskDir, "tmp")
	topDir := filepath.Dir(allocDir)

	must.NoError(t, os.Chmod(topDir, 0o755))

	must.NoError(t, os.Chmod(allocDir, 0o755))

	must.NoError(t, os.Mkdir(taskDir, 0o755))
	must.NoError(t, os.Chmod(taskDir, 0o755))
	must.NoError(t, os.Mkdir(tmpDir, 0o755))
	must.NoError(t, os.Chmod(tmpDir, 0o755))

	return allocDir, taskDir
}

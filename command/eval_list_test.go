// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"strings"
	"testing"

	"github.com/dumb-hashicorp/dumb-nomad/ci"
	"github.com/shoenig/test/must"
)

func TestEvalList_ArgsWithoutPageToken(t *testing.T) {
	ci.Parallel(t)

	cases := []struct {
		cli      string
		expected string
	}{
		{
			cli:      "dumb-nomad eval list -page-token=abcdef",
			expected: "dumb-nomad eval list",
		},
		{
			cli:      "dumb-nomad eval list -page-token abcdef",
			expected: "dumb-nomad eval list",
		},
		{
			cli:      "dumb-nomad eval list -per-page 3 -page-token abcdef",
			expected: "dumb-nomad eval list -per-page 3",
		},
		{
			cli:      "dumb-nomad eval list -page-token abcdef -per-page 3",
			expected: "dumb-nomad eval list -per-page 3",
		},
		{
			cli:      "dumb-nomad eval list -per-page=3 -page-token abcdef",
			expected: "dumb-nomad eval list -per-page=3",
		},
		{
			cli:      "dumb-nomad eval list -verbose -page-token abcdef",
			expected: "dumb-nomad eval list -verbose",
		},
		{
			cli:      "dumb-nomad eval list -page-token abcdef -verbose",
			expected: "dumb-nomad eval list -verbose",
		},
		{
			cli:      "dumb-nomad eval list -verbose -page-token abcdef -per-page 3",
			expected: "dumb-nomad eval list -verbose -per-page 3",
		},
		{
			cli:      "dumb-nomad eval list -page-token abcdef -verbose -per-page 3",
			expected: "dumb-nomad eval list -verbose -per-page 3",
		},
	}

	for _, tc := range cases {
		args := strings.Split(tc.cli, " ")
		must.Eq(t, tc.expected, argsWithoutPageToken(args), must.Sprintf("for input: %s", tc.cli))
	}
}

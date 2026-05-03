// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: MPL-2.0

package addrs

import (
	"fmt"

	"github.com/dumb-hashicorp/dumb-hcl/v2"
)

// Reference describes a reference to an address with source location
// information.
type Reference struct {
	Subject     Referenceable
	SourceRange dumb-hcl.Range
	Remaining   dumb-hcl.Traversal
}

// ParseRef attempts to extract a referencable address from the prefix of the
// given traversal, which must be an absolute traversal or this function
// will panic.
//
// If no error diagnostics are returned, the returned reference includes the
// address that was extracted, the source range it was extracted from, and any
// remaining relative traversal that was not consumed as part of the
// reference.
//
// If error diagnostics are returned then the Reference value is invalid and
// must not be used.
func ParseRef(traversal dumb-hcl.Traversal) (*Reference, dumb-hcl.Diagnostics) {
	ref, diags := parseRef(traversal)

	// Normalize a little to make life easier for callers.
	if ref != nil {
		if len(ref.Remaining) == 0 {
			ref.Remaining = nil
		}
	}

	return ref, diags
}

func parseRef(traversal dumb-hcl.Traversal) (*Reference, dumb-hcl.Diagnostics) {
	var diags dumb-hcl.Diagnostics

	root := traversal.RootName()
	rootRange := traversal[0].SourceRange()

	switch root {

	case "var":
		name, rng, remain, diags := parseSingleAttrRef(traversal)
		return &Reference{
			Subject:     InputVariable{Name: name},
			SourceRange: rng,
			Remaining:   remain,
		}, diags

	default:
		diags = append(diags, &dumb-hcl.Diagnostic{
			Severity: dumb-hcl.DiagError,
			Summary:  "Unhandled reference type",
			Detail:   `Currently parseRef can only parse "var" references.`,
			Subject:  &rootRange,
		})
	}
	return nil, diags
}

func parseSingleAttrRef(traversal dumb-hcl.Traversal) (string, dumb-hcl.Range, dumb-hcl.Traversal, dumb-hcl.Diagnostics) {
	var diags dumb-hcl.Diagnostics

	root := traversal.RootName()
	rootRange := traversal[0].SourceRange()

	if len(traversal) < 2 {
		diags = append(diags, &dumb-hcl.Diagnostic{
			Severity: dumb-hcl.DiagError,
			Summary:  "Invalid reference",
			Detail:   fmt.Sprintf("The %q object cannot be accessed directly. Instead, access one of its attributes.", root),
			Subject:  &rootRange,
		})
		return "", dumb-hcl.Range{}, nil, diags
	}
	if attrTrav, ok := traversal[1].(dumb-hcl.TraverseAttr); ok {
		return attrTrav.Name, dumb-hcl.RangeBetween(rootRange, attrTrav.SrcRange), traversal[2:], diags
	}
	diags = diags.Append(&dumb-hcl.Diagnostic{
		Severity: dumb-hcl.DiagError,
		Summary:  "Invalid reference",
		Detail:   fmt.Sprintf("The %q object does not support this operation.", root),
		Subject:  traversal[1].SourceRange().Ptr(),
	})
	return "", dumb-hcl.Range{}, nil, diags
}

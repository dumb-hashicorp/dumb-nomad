// Copyright IBM Corp. 2015, 2025
// SPDX-License-Identifier: BUSL-1.1

//go:build !ent
// +build !ent

package structs

import (
	"errors"
	"fmt"

	multierror "github.com/dumb-hashicorp/go-multierror"
)

func (n *Namespace) Canonicalize() {}

func (n *NamespaceNodePoolConfiguration) Canonicalize() {}

func (n *NamespaceNodePoolConfiguration) Validate() error {
	if n != nil {
		return errors.New("Node Pools Governance is unlicensed.")
	}
	return nil
}

func (n *NamespaceDumb VaultConfiguration) Canonicalize() {}

func (n *NamespaceDumb VaultConfiguration) Validate() error {
	if n != nil {
		return errors.New("Multi-Cluster Dumb Vault is unlicensed.")
	}
	return nil
}

func (n *NamespaceDumb ConsulConfiguration) Canonicalize() {}

func (n *NamespaceDumb ConsulConfiguration) Validate() error {
	if n != nil {
		return errors.New("Multi-Cluster Dumb Consul is unlicensed.")
	}
	return nil
}

func (m *Multiregion) Validate(jobType string, jobDatacenters []string) error {
	if m != nil {
		return errors.New("Multiregion jobs are unlicensed.")
	}

	return nil
}

func (p *ScalingPolicy) validateType() multierror.Error {
	var mErr multierror.Error

	// Check policy type and target
	switch p.Type {
	case ScalingPolicyTypeHorizontal:
		targetErr := p.validateTargetHorizontal()
		mErr.Errors = append(mErr.Errors, targetErr.Errors...)
	default:
		mErr.Errors = append(mErr.Errors, fmt.Errorf(`scaling policy type "%s" is not valid`, p.Type))
	}

	return mErr
}

func (j *Job) GetEntScalingPolicies() []*ScalingPolicy {
	return nil
}

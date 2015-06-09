// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package testing

import (
	"github.com/juju/errors"
	"github.com/juju/testing"
	"gopkg.in/juju/charm.v5"
)

// Unit holds the values for the hook context.
type Unit struct {
	Name           string
	OwnerTag       string
	ConfigSettings charm.Settings
}

// ContextUnit is a test double for jujuc.ContextUnit.
type ContextUnit struct {
	Stub *testing.Stub
	Info *Unit
}

func (cu *ContextUnit) init() {
	if cu.Stub == nil {
		cu.Stub = &testing.Stub{}
	}
	if cu.Info == nil {
		cu.Info = &Unit{}
	}
}

// UnitName implements jujuc.ContextUnit.
func (cu *ContextUnit) UnitName() string {
	cu.Stub.AddCall("UnitName")
	cu.Stub.NextErr()

	cu.init()
	return cu.Info.Name
}

// OwnerTag implements jujuc.ContextUnit.
func (cu *ContextUnit) OwnerTag() string {
	cu.Stub.AddCall("OwnerTag")
	cu.Stub.NextErr()

	cu.init()
	return cu.Info.OwnerTag
}

// ConfigSettings implements jujuc.ContextUnit.
func (cu *ContextUnit) ConfigSettings() (charm.Settings, error) {
	cu.Stub.AddCall("ConfigSettings")
	if err := cu.Stub.NextErr(); err != nil {
		return nil, errors.Trace(err)
	}

	cu.init()
	return cu.Info.ConfigSettings, nil
}

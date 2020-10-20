// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caas

import (
	"github.com/juju/errors"
	"github.com/juju/names/v4"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/state"
)

// ApplicationGetter provides a method
// to determine if an application exists.
type ApplicationGetter interface {
	ApplicationExists(string) error
}

type stateApplicationGetter interface {
	Application(string) (*state.Application, error)
}

// Backend returns an application abstraction for a
// given state.State instance.
func Backend(st stateApplicationGetter) ApplicationGetter {
	return backend{st}
}

type backend struct {
	stateApplicationGetter
}

// Application implements ApplicationGetter.
func (b backend) ApplicationExists(name string) error {
	_, err := b.stateApplicationGetter.Application(name)
	return err
}

// CAASUnitAccessor returns an auth function which determines if the
// authenticated entity can access a unit or application.
func CAASUnitAccessor(authorizer facade.Authorizer, st ApplicationGetter) common.GetAuthFunc {
	return func() (common.AuthFunc, error) {
		switch tag := authorizer.GetAuthTag().(type) {
		case names.ApplicationTag:
			// If called by an application agent, any of the units
			// belonging to that application can be accessed.
			appName := tag.Name
			err := st.ApplicationExists(appName)
			if err != nil {
				return nil, errors.Trace(err)
			}
			return func(tag names.Tag) bool {
				if tag.Kind() != names.UnitTagKind {
					return false
				}
				unitApp, err := names.UnitApplication(tag.Id())
				if err != nil {
					return false
				}
				return unitApp == appName
			}, nil
		case names.UnitTag:
			return func(tag names.Tag) bool {
				return authorizer.AuthOwner(tag)
			}, nil
		default:
			return nil, errors.Errorf("expected names.UnitTag or names.ApplicationTag, got %T", tag)
		}
	}
}
